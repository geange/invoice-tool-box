package sdk

import "github.com/go-resty/resty/v2"

type SDK struct {
	apiKey    string
	secretKey string

	client *resty.Client
}

func NewSDK(apiKey string, secretKey string) *SDK {
	return &SDK{
		apiKey:    apiKey,
		secretKey: secretKey,
		client:    resty.New(),
	}
}
