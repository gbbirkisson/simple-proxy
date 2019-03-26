package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

/*
	Utilities
*/

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getProxyURL() string {
	return os.Getenv("PROXY_URL")
}

// Get the port to listen on
func getListenAddress() string {
	port := getEnv("PORT", "9900")
	return ":" + port
}

// Log the env variables required for a reverse proxy
func logSetup() {
	log.Printf("Redirecting to url: %s\n", getProxyURL())
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	serveReverseProxy(getProxyURL(), res, req)
}

/*
	Entry
*/

func main() {
	// Log setup values
	logSetup()

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
