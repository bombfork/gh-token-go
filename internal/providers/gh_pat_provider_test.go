package providers

import (
	"testing"
)

func TestGhPatProvider_GetToken_Success(t *testing.T) {
	provider := &ghPatProviderImpl{token: "test-token"}
	token, err := provider.GetToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "test-token" {
		t.Errorf("expected token 'test-token', got %v", token)
	}
}

func TestGhPatProvider_GetToken_Failure(t *testing.T) {
	provider := &ghPatProviderImpl{token: ""}
	_, err := provider.GetToken()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
