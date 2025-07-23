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

	// Modify Response to fix CORS headers
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Del("Access-Control-Allow-Origin")
		resp.Header.Del("Access-Control-Allow-Credentials")
		resp.Header.Del("Access-Control-Allow-Headers")
		resp.Header.Del("Access-Control-Allow-Methods")
		resp.Header.Del("Access-Control-Expose-Headers")
		return nil
	}

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = strings.TrimPrefix(req.URL.Path, pathPrefix)
		req.Host = url.Host
	}

	return proxy
}
