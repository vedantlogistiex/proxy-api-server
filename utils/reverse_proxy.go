package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func CreateReverseProxy(targetHost string, pathPrefix string) *httputil.ReverseProxy {
	targetURL, _ := url.Parse(targetHost)

	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Remove the prefix (e.g., "/v1/tmdb") from the incoming path
			req.URL.Path = strings.TrimPrefix(req.URL.Path, pathPrefix)

			// Set the host and scheme
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host

			// Copy over the original query params
			req.Host = targetURL.Host
		},
	}
}
