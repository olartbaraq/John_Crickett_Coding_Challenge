package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/olartbaraq/load_balancer/utils"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("Received request from %v\n", r.Host)
	// fmt.Printf("%v \n", r.Method)
	// fmt.Printf("%v %v %v \n", r.Host, r.URL, r.Proto)
	// fmt.Printf("%v \n", r.Header["Accept"])
	// fmt.Printf("%v \n", r.Header["User-Agent"])
	// fmt.Printf("Response from server %v %v OK\n", r.Proto, http.StatusOK)
	w.Write([]byte("Hello From Backend Server..."))
}

func NewDevServer(address string) *Server {

	serverUrl, err := url.Parse("http://" + address)
	utils.OnPanicError(err, "invalid server address")

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	return &Server{
		healthCheck: serverUrl.String(),
		Address:     serverUrl,
		ServerMux:   mux,
		Proxy:       ALternateSingleHostReverseProxy(serverUrl),
	}
}

func NewProdServer(healthCheck string, address string) *Server {
	serverUrl, err := url.Parse(address)

	utils.OnPanicError(err, "invalid server address")

	return &Server{
		healthCheck: serverUrl.JoinPath(healthCheck).String(),
		Address:     serverUrl,
		Proxy:       ALternateSingleHostReverseProxy(serverUrl),
	}
}

func (s *Server) IsAlive() bool {

	// randomly choose servers that aren't alive to simulate cases where server goes down.
	if SimulateServerDown() {
		return false
	}

	res, err := http.Get(s.healthCheck)

	if err != nil {
		log.Fatalf("HTTP protocol error occured - %v", err)
		return false
	}

	if res.StatusCode == 200 {
		return true
	}

	return false

}

func (s *Server) Serve(w http.ResponseWriter, req *http.Request) {
	s.Proxy.ServeHTTP(w, req)

}
