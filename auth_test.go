package epcc_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/moltin/go-epcc-api-client"
	"github.com/stretchr/testify/assert"
)

func fakeHandleAuth(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(500)
		return
	}
	bodyStr := string(body)

	switch {
	case req.Method == "POST" && bodyStr == "client_id=validClientID&client_secret=validClientSecret&grant_type=client_credentials":
		responseJSON := `{` +
			`"expires":1598636721,` +
			`"access_token":"f64e7f07b10f710a15e4f41d670f0d7d7d4e415d",` +
			`"identifier":"client_credentials",` +
			`"expires_in":3600,` +
			`"token_type":"Bearer"` +
			`}`
		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))

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
		{"invalidClientID", "invalidClientSecret", "", errors.New("authentication error")},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleAuth))
	limitTimeout := 10 * time.Millisecond
	clientTimeout := 10 * time.Second
	client := epcc.NewClient(testServer.URL, limitTimeout, clientTimeout)

	for _, test := range tests {

		os.Setenv("GO_EPCC_CLIENT_ID", test.clientID)
		os.Setenv("GO_EPCC_CLIENT_SECRET", test.clientSecret)

		token, err := epcc.Auth(*client)
		assert.Equal(t, test.expectedToken, token)
		assert.Equal(t, test.err, err)
	}
}
