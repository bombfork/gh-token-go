package providers

type ghPatProviderImpl struct {
	token string
}

func NewGhPatProvider(pat string) (*ghPatProviderImpl, error) {
	return &ghPatProviderImpl{
		token: pat,
	}, nil
}

func (t *ghPatProviderImpl) GetToken() (string, error) {
	return t.token, nil
}
