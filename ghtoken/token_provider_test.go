package ghtoken

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bombfork/gh-token-go/internal/testutils"
)

func TestNewGhTokenProviderDefault_NoCreds_Failure(t *testing.T) {
	if err := os.Unsetenv("GH_TKN"); err != nil {
		t.Fatalf("failed to unset GH_TKN: %v", err)
	}
	if err := os.Unsetenv("GH_TOKEN"); err != nil {
		t.Fatalf("failed to unset GH_TOKEN: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_INST_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_PRIVATE_KEY"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Unsetenv("GITHUB_TOKEN"); err != nil {
		t.Fatalf("failed to unset GITHUB_TOKEN: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if provider != nil {
		t.Fatalf("expected provider to be nil, got %v", provider)
	}
}

func TestNewGhTokenProviderDefault_Using_PAT_Success(t *testing.T) {
	if err := os.Unsetenv("GH_TOKEN"); err != nil {
		t.Fatalf("failed to unset GH_TOKEN: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_INST_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_PRIVATE_KEY"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Unsetenv("GITHUB_TOKEN"); err != nil {
		t.Fatalf("failed to unset GITHUB_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TKN", "mocked-token-value"); err != nil {
		t.Fatalf("failed to set GH_TKN: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestNewGhTokenProviderDefault_Using_Std_PAT_Success(t *testing.T) {
	if err := os.Unsetenv("GH_TKN"); err != nil {
		t.Fatalf("failed to unset GH_TKN: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_INST_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_PRIVATE_KEY"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Unsetenv("GH_TOKEN"); err != nil {
		t.Fatalf("failed to unset GITHUB_TOKEN: %v", err)
	}
	if err := os.Setenv("GITHUB_TOKEN", "mocked-token-value"); err != nil {
		t.Fatalf("failed to set GITHUB_TOKEN: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestNewGhTokenProviderDefault_Using_Alt_Std_PAT_Success(t *testing.T) {
	if err := os.Unsetenv("GH_TKN"); err != nil {
		t.Fatalf("failed to unset GH_TKN: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_INST_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_PRIVATE_KEY"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Unsetenv("GITHUB_TOKEN"); err != nil {
		t.Fatalf("failed to unset GITHUB_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TOKEN", "mocked-token-value"); err != nil {
		t.Fatalf("failed to set GH_TOKEN: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestNewGhTokenProviderDefault_Using_PAT_and_Std_PAT_Success(t *testing.T) {
	if err := os.Unsetenv("GH_TKN_APP_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_INST_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_PRIVATE_KEY"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Unsetenv("GH_TOKEN"); err != nil {
		t.Fatalf("failed to unset GITHUB_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TKN", "mocked-token-value"); err != nil {
		t.Fatalf("failed to set GH_TKN: %v", err)
	}
	if err := os.Setenv("GITHUB_TOKEN", "should-not-be-token-value"); err != nil {
		t.Fatalf("failed to set GITHUB_TOKEN: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestNewGhTokenProviderDefault_Using_Std_PAT_and_Alt_Std_PAT_Success(t *testing.T) {
	if err := os.Unsetenv("GH_TKN"); err != nil {
		t.Fatalf("failed to unset GH_TKN: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_INST_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_PRIVATE_KEY"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Setenv("GITHUB_TOKEN", "mocked-token-value"); err != nil {
		t.Fatalf("failed to set GITHUB_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TOKEN", "should-not-be-token-value"); err != nil {
		t.Fatalf("failed to set GH_TOKEN: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestNewGhTokenProviderDefault_Using_All_PAT_Success(t *testing.T) {
	if err := os.Unsetenv("GH_TKN_APP_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_INST_ID"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Unsetenv("GH_TKN_APP_PRIVATE_KEY"); err != nil {
		t.Fatalf("failed to unset GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Setenv("GH_TKN", "mocked-token-value"); err != nil {
		t.Fatalf("failed to set GH_TKN: %v", err)
	}
	if err := os.Setenv("GITHUB_TOKEN", "should-not-be-used-token-value"); err != nil {
		t.Fatalf("failed to set GITHUB_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TOKEN", "should-not-be-token-value"); err != nil {
		t.Fatalf("failed to set GH_TOKEN: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestNewGhTokenProviderDefault_Using_Gh_App_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte(`{"token":"mocked-token-value"}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer testServer.Close()

	if err := os.Setenv("GH_TKN_APP_ID", "12345"); err != nil {
		t.Fatalf("failed to set GH_TKN_APP_ID: %v", err)
	}
	if err := os.Setenv("GH_TKN_APP_INST_ID", "67890"); err != nil {
		t.Fatalf("failed to set GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Setenv("GH_TKN_APP_PRIVATE_KEY", testutils.FakeValidPemKey); err != nil {
		t.Fatalf("failed to set GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Unsetenv("GH_TKN"); err != nil {
		t.Fatalf("failed to unset GH_TKN: %v", err)
	}
	if err := os.Unsetenv("GITHUB_TOKEN"); err != nil {
		t.Fatalf("failed to unset GITHUB_TOKEN: %v", err)
	}
	if err := os.Unsetenv("GH_TOKEN"); err != nil {
		t.Fatalf("failed to unset GH_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TKN_API_URL", testServer.URL); err != nil {
		t.Fatalf("failed to set GH_TKN_API_URL: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}

func TestNewGhTokenProviderDefault_Using_All_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte(`{"token":"mocked-token-value"}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	}))
	defer testServer.Close()

	if err := os.Setenv("GH_TKN_APP_ID", "12345"); err != nil {
		t.Fatalf("failed to set GH_TKN_APP_ID: %v", err)
	}
	if err := os.Setenv("GH_TKN_APP_INST_ID", "67890"); err != nil {
		t.Fatalf("failed to set GH_TKN_APP_INST_ID: %v", err)
	}
	if err := os.Setenv("GH_TKN_APP_PRIVATE_KEY", testutils.FakeValidPemKey); err != nil {
		t.Fatalf("failed to set GH_TKN_APP_PRIVATE_KEY: %v", err)
	}
	if err := os.Setenv("GH_TKN", "should-not-be-used-token-value"); err != nil {
		t.Fatalf("failed to set GH_TKN: %v", err)
	}
	if err := os.Setenv("GITHUB_TOKEN", "should-not-be-used-token-value"); err != nil {
		t.Fatalf("failed to set GITHUB_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TOKEN", "should-not-be-used-token-value"); err != nil {
		t.Fatalf("failed to set GH_TOKEN: %v", err)
	}
	if err := os.Setenv("GH_TKN_API_URL", testServer.URL); err != nil {
		t.Fatalf("failed to set GH_TKN_API_URL: %v", err)
	}
	provider, err := NewGhTokenProviderDefault()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider not be nil")
	}

	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "mocked-token-value" {
		t.Fatalf("expected token 'mocked-token-value', got %v", token)
	}
}
