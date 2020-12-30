package url

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Fetch(){
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
/*
a goroutine is a concurrent function execution
a channel is a communication mechanism that allows one goroutine to pass values of a specified type to another go routine
when one go routine attempts a send or receive on a channel, it blocks until another goroutine attempts the corresponding receive or send operation

 */

func ConcurrentFetch() {
	start := time.Now()
	ch := make(chan string) // creating a channel of strings
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a go routine
	}
	for range os.Args[1:] { //receives and prints those lines
		fmt.Println(<-ch) //receives all of them
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // Reads the body of the response and discards it, returns the byte count
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v",url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
