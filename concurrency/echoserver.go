package concurrency

import (
	"io"
	"log"
	"net"
)

// Clock1 is a TCP server that periodically writes the time.

func EchoServer(){
	listener, err := net.Listen("tcp","localhost:7111")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnEcho(conn)
	}
}

func handleConnEcho(c net.Conn){
	defer c.Close()
	io.Copy(c,c)
}
