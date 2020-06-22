package testutils

import (
	"net/http"
	"strings"
)

const AuthTokenHeader = "X-Auth-Token"

func TableTestName(desc string) string {
	return strings.Replace(desc, " ", "_", -1)
}

func SetAuthTokenHeader(req *http.Request, token string) {
	req.Header.Add(AuthTokenHeader, token)
}
