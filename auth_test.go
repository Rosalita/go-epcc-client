package epcc_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/moltin/go-epcc-client"
	"github.com/stretchr/testify/assert"
	"bytes"
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

func TestAuth(t *testing.T) {

	tests := []struct {
		clientID      string
		clientSecret  string
		expectedToken string
		err           error
	}{
		{"validClientID", "validClientSecret", "f64e7f07b10f710a15e4f41d670f0d7d7d4e415d", nil},
		{"invalidClientID", "invalidClientSecret", "", errors.New("error: unexpected status 403 Forbidden")},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleAuth))
	limitTimeout := 10 * time.Millisecond
	clientTimeout := 10 * time.Second
	client := epcc.NewClient(&testServer.URL, limitTimeout, clientTimeout)

	for _, test := range tests {

		epcc.Cfg.Credentials.ClientID = test.clientID
		epcc.Cfg.Credentials.ClientSecret = test.clientSecret

		token, err := epcc.Auth(*client)
		assert.Equal(t, test.expectedToken, token)
		assert.Equal(t, test.err, err)
	}
}
