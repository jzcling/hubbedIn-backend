package interfaces

import "net/http"

// HTTPClient describes a default http client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
