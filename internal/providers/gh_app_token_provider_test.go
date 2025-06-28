package providers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bombfork/gh-token-go/internal/testutils"
)

func TestGhAppTokenProvider_GetToken_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte(`{"token":"mocked-token-value"}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer testServer.Close()

	provider := &ghAppTokenProviderImpl{
		pemKey:         testutils.FakeValidPemKey,
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
		if _, err := w.Write([]byte(`{"token":"mocked-token-value"}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
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
		if _, err := w.Write([]byte(`{"error":"Forbidden"}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer testServer.Close()

	provider := &ghAppTokenProviderImpl{
		pemKey:         testutils.FakeValidPemKey,
		appID:          12345,
		installationID: 67890,
		ghApiUrl:       testServer.URL,
	}

	_, err := provider.GetToken()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
