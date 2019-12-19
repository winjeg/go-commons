package openapi

type DefaultProvider struct {
	AppKey    string
	AppSecret string
}

func (dp DefaultProvider) GetSecret(key string) (string, error) {
	return dp.AppSecret, nil
}
