package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ghAppTokenProviderImpl struct {
	pemKey         string
	appID          int
	installationID int
	token          string
	expiresAt      time.Time
	ghApiUrl       string
}

func NewGhAppTokenProvider(pemKey string, appID int, installationID int, apiUrl string) (*ghAppTokenProviderImpl, error) {
	log.Println("Creating a new GhAppTokenProvider")
	p := &ghAppTokenProviderImpl{
		pemKey:         pemKey,
		appID:          appID,
		installationID: installationID,
		ghApiUrl:       apiUrl,
	}
	if err := p.refreshToken(); err != nil {
		return nil, fmt.Errorf("could not create GhAppTokenProvider with the provided parameters: appId=%d, installationId=%d, pemKey=%s, apiUrl=%s, err=%v", appID, installationID, pemKey, apiUrl, err)
	}
	log.Println("GhAppTokenProvider created successfully")
	return p, nil
}

func (t *ghAppTokenProviderImpl) GetToken() (string, error) {
	if t.token == "" || time.Now().After(t.expiresAt) {
		if err := t.refreshToken(); err != nil {
			return "", fmt.Errorf("could not refresh GH APP token: %w", err)
		}
	}
	return t.token, nil
}

func (t *ghAppTokenProviderImpl) refreshToken() error {
	log.Println("Refreshing GitHub App token")
	expirestAt := time.Now().Add(time.Minute * 10)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": expirestAt.Unix(),
		"iss": t.appID,
	})

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(t.pemKey))
	if err != nil {
		return err
	}

	signedToken, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/app/installations/%d/access_tokens", t.ghApiUrl, t.installationID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+signedToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to get installation token: %s", body)
	}

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	t.token = result["token"].(string)
	t.expiresAt = expirestAt
	log.Printf("GitHub App token refreshed successfully, expires at: %s", t.expiresAt)
	return nil
}
