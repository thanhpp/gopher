package vtclient

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type basicAuth struct {
	username string
	password string
}

func newBasicAuth(username, password string) *basicAuth {
	return &basicAuth{
		username: username,
		password: password,
	}
}

func (b *basicAuth) genEncodedString() string {
	auth := fmt.Sprintf("%s:%s", b.username, b.password)

	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (b *basicAuth) addHeader(r *http.Request) {
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", b.genEncodedString()))
}
