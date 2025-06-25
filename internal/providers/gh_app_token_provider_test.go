package providers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	fakeValidPemKey = `-----BEGIN PRIVATE KEY-----
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
N8vp4wZ9cdziWz+R4j0JWhyHiNxa
-----END PRIVATE KEY-----`
)

func TestGhAppTokenProvider_GetToken_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"token":"mocked-token-value"}`))
	}))
	defer testServer.Close()

	provider := &ghAppTokenProviderImpl{
		pemKey:         fakeValidPemKey,
		appID:          12345,
		installationID: 67890,
		ghApiUrl:       testServer.URL,
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Errorf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestGhAppTokenProvider_GetToken_FailureInvalidPemKey(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"token":"mocked-token-value"}`))
	}))
	defer testServer.Close()

	provider := &ghAppTokenProviderImpl{
		pemKey:         "invalid-pem-key",
		appID:          12345,
		installationID: 67890,
		ghApiUrl:       testServer.URL,
	}

	_, err := provider.GetToken()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGhAppTokenProvider_GetToken_FailureForbidden(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"error":"Forbidden"}`))
	}))
	defer testServer.Close()

	provider := &ghAppTokenProviderImpl{
		pemKey:         fakeValidPemKey,
		appID:          12345,
		installationID: 67890,
		ghApiUrl:       testServer.URL,
	}

	_, err := provider.GetToken()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
