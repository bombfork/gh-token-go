package ghtoken

import "testing"

func TestDefaultGhTokenProviderCfg_Values(t *testing.T) {
	cfg := DefaultGhTokenProviderCfg

	if cfg.AppIdVarName != "GH_TKN_APP_ID" {
		t.Errorf("AppIdVarName: got %q, want %q", cfg.AppIdVarName, "GH_TKN_APP_ID")
	}
	if cfg.AppInstIdVarName != "GH_TKN_APP_INST_ID" {
		t.Errorf("AppInstIdVarName: got %q, want %q", cfg.AppInstIdVarName, "GH_TKN_APP_INST_ID")
	}
	if cfg.AppPemKeyVarName != "GH_TKN_APP_PRIVATE_KEY" {
		t.Errorf("AppPemKeyVarName: got %q, want %q", cfg.AppPemKeyVarName, "GH_TKN_APP_PRIVATE_KEY")
	}
	if cfg.PatVarName != "GH_TKN" {
		t.Errorf("PatVarName: got %q, want %q", cfg.PatVarName, "GH_TKN")
	}
	if cfg.GhApiUrlVarName != "GH_TKN_API_URL" {
		t.Errorf("GhApiUrlVarName: got %q, want %q", cfg.GhApiUrlVarName, "GH_TKN_API_URL")
	}
}

func TestGhTokenProviderCfg_CustomValues(t *testing.T) {
	cfg := GhTokenProviderCfg{
		AppIdVarName:     "APP_ID",
		AppInstIdVarName: "APP_INST_ID",
		AppPemKeyVarName: "PRIVATE_KEY",
		PatVarName:       "PAT",
		GhApiUrlVarName:  "API_URL",
	}

	if cfg.AppIdVarName != "APP_ID" {
		t.Errorf("AppIdVarName: got %q, want %q", cfg.AppIdVarName, "APP_ID")
	}
	if cfg.AppInstIdVarName != "APP_INST_ID" {
		t.Errorf("AppInstIdVarName: got %q, want %q", cfg.AppInstIdVarName, "APP_INST_ID")
	}
	if cfg.AppPemKeyVarName != "PRIVATE_KEY" {
		t.Errorf("AppPemKeyVarName: got %q, want %q", cfg.AppPemKeyVarName, "PRIVATE_KEY")
	}
	if cfg.PatVarName != "PAT" {
		t.Errorf("PatVarName: got %q, want %q", cfg.PatVarName, "PAT")
	}
	if cfg.GhApiUrlVarName != "API_URL" {
		t.Errorf("GhApiUrlVarName: got %q, want %q", cfg.GhApiUrlVarName, "API_URL")
	}
}
