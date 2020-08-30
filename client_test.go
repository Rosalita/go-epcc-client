package epcc

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func fakeHandleAuth(rw http.ResponseWriter, req *http.Request) {
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(req.Body)
	if err != nil {
		rw.WriteHeader(500)
		return
	}

	switch {
	case req.URL.String() == "/oauth/access_token" && req.Method == "POST" && buffer.String() == "client_id=validClientID&client_secret=validClientSecret&grant_type=client_credentials":
		responseJSON := `{` +
			`"expires":1598636721,` +
			`"access_token":"f64e7f07b10f710a15e4f41d670f0d7d7d4e415d",` +
			`"identifier":"client_credentials",` +
			`"expires_in":3600,` +
			`"token_type":"Bearer"` +
			`}`
		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))
	case req.URL.String() == "/oauth/access_token" && req.Method == "POST" && buffer.String() == "client_id=invalidClientID&client_secret=invalidClientSecret&grant_type=client_credentials":
		rw.WriteHeader(403)
	default:
		rw.WriteHeader(500)
	}
}
func TestAuthenticate(t *testing.T) {

	limitTimeout := 10 * time.Millisecond
	clientTimeout := 10 * time.Second
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleAuth))

	tests := []struct {
		clientID            string
		clientSecret        string
		expectedAccessToken string
		err                 error
	}{
		{"validClientID", "validClientSecret", "f64e7f07b10f710a15e4f41d670f0d7d7d4e415d", nil},
		{"invalidClientID", "invalidClientSecret", "", errors.New("error: unexpected status 403 Forbidden")},
	}

	for _, test := range tests {

		newClient := NewClient(&testServer.URL, limitTimeout, clientTimeout)

		Cfg.Credentials.ClientID = test.clientID
		Cfg.Credentials.ClientSecret = test.clientSecret
		assert.Equal(t, "", newClient.accessToken)
		err := newClient.Authenticate()
		assert.Equal(t, test.expectedAccessToken, newClient.accessToken)
		assert.Equal(t, test.err, err)
	}
}
