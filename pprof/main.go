package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	var mu1, mu2 sync.Mutex
	mu1.Lock()
	mu2.Lock()
}

func main() {
	http.HandleFunc("/", hiHandler)
	log.Fatal(http.ListenAndServe(":6060", nil))
}
