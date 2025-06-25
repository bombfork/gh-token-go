package ghtoken

type ghPatProviderImpl struct {
	token string
}

func newGhPatProvider(pat string) (GhTokenProvider, error) {
	return &ghPatProviderImpl{
		token: pat,
	}, nil
}

func (t *ghPatProviderImpl) GetToken() (string, error) {
	return t.token, nil
}
