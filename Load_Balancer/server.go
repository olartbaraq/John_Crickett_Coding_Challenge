package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Server struct {
	healthCheck string
	Address     *url.URL
	ServerMux   http.Handler
	Proxy       *httputil.ReverseProxy
}

func SimulateServerDown() bool {
	ArrayOfTrueFalse := []bool{true, false}
	// pick a boolean value at random
	index := rand.Intn(len(ArrayOfTrueFalse))
	return ArrayOfTrueFalse[index]
}

func JoinToSingleSlash(firstUrl, secondUrl string) string {
	firstSlash := strings.HasSuffix(firstUrl, "/")
	secondSlash := strings.HasPrefix(secondUrl, "/")

	switch {
	case firstSlash && secondSlash:
		return firstUrl + secondUrl[1:] // let secondUrl start from index 1 to the end

	case !firstSlash && !secondSlash:
		return fmt.Sprintf("%v/%v", firstUrl, secondUrl)
	}
	return firstUrl + secondUrl
}

func ALternateSingleHostReverseProxy(url *url.URL) *httputil.ReverseProxy {
	urlQuery := url.RawQuery

	director := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.Host = url.Host
		req.URL.Path = JoinToSingleSlash(url.Path, req.URL.Path)
		if urlQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = urlQuery + req.URL.RawQuery
		}
		req.URL.RawQuery = urlQuery + "&" + req.URL.RawQuery
	}
	return &httputil.ReverseProxy{Director: director}
}
