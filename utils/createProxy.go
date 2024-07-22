package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func CreateReverseProxy(target string, pathPrefix string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = strings.TrimPrefix(req.URL.Path, pathPrefix)
		req.Host = url.Host
	}
	return proxy
}
