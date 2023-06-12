package utils

import (
	"net/http"
	"strings"
)

func CheckIfClientAcceptEncoding(r *http.Request, encoding string) bool {
	if r == nil {
		return false
	}
	acceptEncodingHeader, ok := r.Header["Accept-Encoding"]
	return ok && len(acceptEncodingHeader) > 0 && strings.Contains(acceptEncodingHeader[0], encoding)
}
