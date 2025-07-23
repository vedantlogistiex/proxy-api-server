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
		// Remove existing CORS headers
		resp.Header.Del("Access-Control-Allow-Origin")
		resp.Header.Del("Access-Control-Allow-Credentials")
		resp.Header.Del("Access-Control-Allow-Headers")
		resp.Header.Del("Access-Control-Allow-Methods")
		resp.Header.Del("Access-Control-Expose-Headers")

		// You can set Access-Control-Allow-Origin here OR let Gin middleware set it
		// If you want to handle dynamic Origin properly, you can skip setting it here
		return nil
	}

	// Director to rewrite the URL path
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = strings.TrimPrefix(req.URL.Path, pathPrefix)
		req.Host = url.Host
	}

	return proxy
}
