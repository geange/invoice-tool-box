package sdk

import (
	"encoding/json"
)

func (s *SDK) Token() (*Token, error) {
	host := "https://aip.baidubce.com/oauth/2.0/token"
	resp, err := s.client.R().SetQueryParams(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     s.apiKey,
		"client_secret": s.secretKey,
	}).Get(host)
	if err != nil {
		return nil, err
	}

	var response Token
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type Token struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
}
