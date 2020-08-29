package epcc

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"errors"
)

type authResponse struct {
	Expires     int    `json:"expires"`
	ExpiresIn   int    `json:"expires_in"`
	Identifier  string `json:"identifier"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}



//Auth returns an AccessToken or an Error
func Auth(client Client)(string, error) {

	// TODO check env vars aren't nil
	reqURL, err := url.Parse(client.BaseURL)

	values := url.Values{}
	values.Set("client_id", os.Getenv("GO_EPCC_CLIENT_ID"))
	values.Set("client_secret", os.Getenv("GO_EPCC_CLIENT_SECRET"))
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

	var buffer bytes.Buffer
	buffer.ReadFrom(res.Body)

	if buffer.String() == ""{
		return "", errors.New("authentication error")
	}

	var authResponse authResponse
	if err := json.Unmarshal(buffer.Bytes(), &authResponse); err != nil {
		return "", err
	}

	log.Println("authentication successful")
	return authResponse.AccessToken, nil
}
