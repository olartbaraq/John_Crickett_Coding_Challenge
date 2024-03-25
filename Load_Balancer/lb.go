package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type LoadBalancer struct {
	Port     int
	LastPort int
	Count    int
	Servers  []*Server
	Config   LoadBalancerConfig
}

func (lb *LoadBalancer) getNextServer() *Server {

	lb.Count++
	//fmt.Printf("initial count is %v", lb.Count)
	server := lb.Servers[lb.Count%len(lb.Servers)]

	if !server.IsAlive() {
		lb.Count++
		//fmt.Printf("recurssion count is %v", lb.Count)
		lb.getNextServer()
	}

	return server
}

func portInuse(port int) bool {
	_, err := http.Get("http://localhost:" + strconv.Itoa(port))
	return err == nil
}

func (lb *LoadBalancer) getNextPort() int {

	lb.LastPort++
	//fmt.Printf("initial port is %v", lb.LastPort)

	for {
		if portInuse(lb.LastPort) {
			lb.LastPort++
			//fmt.Printf("loop port is %v", lb.LastPort)
		} else {
			return lb.LastPort
		}
	}
}

func (lb *LoadBalancer) Serve(w http.ResponseWriter, r *http.Request) {
	server := lb.getNextServer()

	fmt.Printf("Sending request to server %v\n", server.Address)

	server.Serve(w, r)

}

func (lb *LoadBalancer) StartDemoServers(wg *sync.WaitGroup) {
	for _, server := range lb.Servers {
		wg.Add(1)

		go func(server *Server) {
			defer wg.Done()
			http.ListenAndServe(server.Address.Host, server.ServerMux)
		}(server)
	}

}

func (lb *LoadBalancer) StartLB(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		http.HandleFunc("/", lb.Serve)
		http.ListenAndServe(":"+strconv.Itoa(lb.Port), nil)
	}()

}

func (lb *LoadBalancer) Start() {
	wg := sync.WaitGroup{}

	defer wg.Wait()

	if lb.Config.Env == "dev" {
		lb.StartDemoServers(&wg)
		lb.StartLB(&wg)

	} else {
		lb.StartLB(&wg)
	}

}
