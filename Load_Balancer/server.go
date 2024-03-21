package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	healthCheck string
	Address     *url.URL
	ServerMux   http.Handler
	Proxy       *httputil.ReverseProxy
}
