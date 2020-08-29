package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type authResponse struct {
	Expires     int    `json:"expires"`
	ExpiresIn   int    `json:"expires_in"`
	Identifier  string `json:"identifier"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

//Auth returns an AccessToken or an Error
func Auth(client Client) (string, error) {

	reqURL, err := url.Parse(client.BaseURL)

	reqURL.Path = fmt.Sprintf("/oauth/access_token")

	values := url.Values{}
	values.Set("client_id", Cfg.Credentials.ClientID)
	values.Set("client_secret", Cfg.Credentials.ClientSecret)
	values.Set("grant_type", "client_credentials")

	body := strings.NewReader(values.Encode())

	req, err := http.NewRequest("POST", reqURL.String(), body)
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("error: unexpected status %s", res.Status)
	}

	var buffer bytes.Buffer
	buffer.ReadFrom(res.Body)

	var authResponse authResponse
	if err := json.Unmarshal(buffer.Bytes(), &authResponse); err != nil {
		return "", err
	}

	log.Println("authentication successful")
	return authResponse.AccessToken, nil
}
