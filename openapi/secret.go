package openapi

// the interface to get the secret
type SecretKeeper interface {
	GetSecret(key string) (string, error)
}

const (
	EmptyString = ""
)
