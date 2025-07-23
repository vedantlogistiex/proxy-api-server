func CreateReverseProxy(targetHost string, prefix string) *httputil.ReverseProxy {
	url, _ := url.Parse(targetHost)
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Modify the response to fix CORS
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Set("Access-Control-Allow-Origin", "*")  // Force single origin
		resp.Header.Del("Access-Control-Allow-Credentials")
		resp.Header.Del("Access-Control-Allow-Headers")
		resp.Header.Del("Access-Control-Allow-Methods")
		resp.Header.Del("Access-Control-Expose-Headers")
		return nil
	}

	proxy.Director = func(req *http.Request) {
		// existing director logic...
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
		req.Host = url.Host
	}

	return proxy
}
