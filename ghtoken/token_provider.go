package ghtoken

import (
	"os"
	"strconv"

	"github.com/bombfork/gh-token-go/internal/providers"
)

type ErrNoCredsProvided struct{}

func (e ErrNoCredsProvided) Error() string {
	return "no credentials provided. check the docs."
}

type GhTokenProvider interface {
	GetToken() (string, error)
}

func NewGhTokenProviderDefault() (GhTokenProvider, error) {
	return NewGhTokenProvider(DefaultGhTokenProviderCfg)
}

func NewGhTokenProvider(cfg GhTokenProviderCfg) (GhTokenProvider, error) {
	appIdStr, isSetAppId := os.LookupEnv(cfg.AppIdVarName)
	appInstIdStr, isSetAppInstId := os.LookupEnv(cfg.AppInstIdVarName)
	appPemKey, isSetAppPrivKey := os.LookupEnv(cfg.AppPemKeyVarName)
	patVal, isSetPatVar := os.LookupEnv(cfg.PatVarName)
	ghApiUrl, isSetGhApiUrl := os.LookupEnv(cfg.GhApiUrlVarName)
	stdPatVal, isSetStdPatVar := os.LookupEnv("GITHUB_TOKEN")
	stdPatAltVal, isSetStdPatAltVar := os.LookupEnv("GH_TOKEN")

	if !isSetGhApiUrl {
		ghApiUrl = "https://api.github.com"
	}

	if isSetAppId && isSetAppInstId && isSetAppPrivKey {
		appId, err := strconv.Atoi(appIdStr)
		if err != nil {
			return nil, err
		}
		appInstId, err := strconv.Atoi(appInstIdStr)
		if err != nil {
			return nil, err
		}
		return providers.NewGhAppTokenProvider(appPemKey, appId, appInstId, ghApiUrl)
	} else if isSetPatVar {
		return providers.NewGhPatProvider(patVal)
	} else if isSetStdPatVar {
		return providers.NewGhPatProvider(stdPatVal)
	} else if isSetStdPatAltVar {
		return providers.NewGhPatProvider(stdPatAltVal)
	} else {
		return nil, ErrNoCredsProvided{}
	}
}
