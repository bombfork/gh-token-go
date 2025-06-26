package ghtoken

import (
	"os"
	"strconv"

	"github.com/bombfork/gh-token-go/internal/providers"
)

type ErrNoCredsProvided struct{}

func (e ErrNoCredsProvided) Error() string {
	return "No credentials provided. Set either GH_TKN_APP_ID, GH_TKN_APP_INST_ID and GH_TKN_APP_PRIVATE_KEY, GH_TKN, GITHUB_TOKEN or GH_TOKEN environment variables."
}

type GhTokenProvider interface {
	GetToken() (string, error)
}

func NewGhTokenProvider() (GhTokenProvider, error) {
	appIdStr, isSetAppId := os.LookupEnv("GH_TKN_APP_ID")
	appInstIdStr, isSetAppInstId := os.LookupEnv("GH_TKN_APP_INST_ID")
	appPemKey, isSetAppPrivKey := os.LookupEnv("GH_TKN_APP_PRIVATE_KEY")
	patVal, isSetPatVar := os.LookupEnv("GH_TKN")
	stdPatVal, isSetStdPatVar := os.LookupEnv("GITHUB_TOKEN")
	stdPatAltVal, isSetStdPatAltVar := os.LookupEnv("GH_TOKEN")
	ghApiUrl, isSetGhApiUrl := os.LookupEnv("GH_TKN_API_URL")
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
