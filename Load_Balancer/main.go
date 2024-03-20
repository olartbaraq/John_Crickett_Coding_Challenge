package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %v \n", r.Host)
	fmt.Printf("%v \n", r.Method)
	fmt.Printf("%v %v %v \n", r.Host, r.URL, r.Proto)
	fmt.Printf("%v \n", r.Header["Accept"])
	fmt.Printf("%v \n", r.Header["User-Agent"])
	fmt.Printf("Replied with a hello message")
	//w.Write([]byte("This is my home page"))
}

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go loadBalancer(&wg)

	http.HandleFunc("/", getHome)

	log.Fatal(http.ListenAndServe(":8000", nil))

	wg.Wait()
}
