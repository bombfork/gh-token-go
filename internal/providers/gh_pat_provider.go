package providers

type ghPatProviderImpl struct {
	token string
}

type ErrEmptyToken struct{}

func (e ErrEmptyToken) Error() string {
	return "GitHub PAT is empty"
}

func NewGhPatProvider(pat string) (*ghPatProviderImpl, error) {
	return &ghPatProviderImpl{
		token: pat,
	}, nil
}

func (t *ghPatProviderImpl) GetToken() (string, error) {
	if t.token == "" {
		return "", ErrEmptyToken{}
	}
	return t.token, nil
}
