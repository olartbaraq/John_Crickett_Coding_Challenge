package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %v \n", r.Host)
	fmt.Printf("%v \n", r.Method)
	fmt.Printf("%v %v %v \n", r.Host, r.URL, r.Proto)
	fmt.Printf("%v \n", r.Header["Accept"])
	fmt.Printf("%v \n", r.Header["User-Agent"])
	fmt.Printf("Response from server %v %v OK\n", r.Proto, http.StatusOK)
	w.Write([]byte("Hello From Backend Server"))
}

func loadBalancer(wg *sync.WaitGroup) {

	http.HandleFunc("/", getRoot)
	//http.HandleFunc("/hello", getHello)

	wg.Done()
	log.Fatal(http.ListenAndServe(":80", nil))
}
