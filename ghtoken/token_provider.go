package ghtoken

import (
	"os"
	"strconv"
)

type ErrNoCredsProvided struct{}

func (e ErrNoCredsProvided) Error() string {
	return "No credentials provided. Set either GH_TOKEN_APP_ID, GH_TOKEN_APP_INST_ID and GH_TOKEN_APP_PRIVATE_KEY or GH_TOKEN or GITHUB_TOKEN environment variables."
}

type GhTokenProvider interface {
	GetToken() (string, error)
}

func NewGhTokenProvider() (GhTokenProvider, error) {
	appIdStr, isSetAppId := os.LookupEnv("GH_TOKEN_APP_ID")
	appInstIdStr, isSetAppInstId := os.LookupEnv("GH_TOKEN_APP_INST_ID")
	appPemKey, isSetAppPrivKey := os.LookupEnv("GH_TOKEN_APP_PRIVATE_KEY")
	patVal, isSetPatVar := os.LookupEnv("GH_TOKEN")
	stdPatVal, isSetStdPatVar := os.LookupEnv("GITHUB_TOKEN")

	if isSetAppId && isSetAppInstId && isSetAppPrivKey {
		appId, err := strconv.Atoi(appIdStr)
		if err != nil {
			return nil, err
		}
		appInstId, err := strconv.Atoi(appInstIdStr)
		if err != nil {
			return nil, err
		}
		return newGhAppTokenProvider(appPemKey, appId, appInstId)
	} else if isSetPatVar {
		return newGhPatTokenProvider(patVal)
	} else if isSetStdPatVar {
		return newGhPatTokenProvider(stdPatVal)
	} else {
		return nil, ErrNoCredsProvided{}
	}
}
