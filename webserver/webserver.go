package webserver

import (
	"fmt"
	"gobook/functions"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

/*
the server runs the handler for each incoming request in a separate go routine so that it can serve multiple request simultaneously
if two concurrent request try to update count at the same time, it might not be incremented consistently. the program would have a serious bug called race condition
con un mutex, solo una go routine puede acceder a la variable a la vez
 */

func WebServer(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/count",counter)
	http.HandleFunc("/draw",functions.Draw)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r * http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
