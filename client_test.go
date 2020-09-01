package epcc

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {

		testServer := httptest.NewServer(http.HandlerFunc(fakeHandleAuth))
		options := ClientOptions{
			BaseURL: testServer.URL,
			ClientTimeout: 10 * time.Second,
			RetryLimitTimeout: 10 * time.Millisecond,
		}

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

		cfg.Credentials.ClientID = test.clientID
		cfg.Credentials.ClientSecret = test.clientSecret

		client := NewClient(options)
		assert.Equal(t, "", client.accessToken)
		err := client.Authenticate()
		assert.Equal(t, test.expectedAccessToken, client.accessToken)
		assert.Equal(t, test.err, err)
	}
}
