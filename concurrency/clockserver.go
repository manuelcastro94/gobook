package concurrency

import (
	"io"
	"log"
	"net"
	"time"
)

// Clock1 is a TCP server that periodically writes the time.

func Clock(){
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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	for {
		_,err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1*time.Second)
	}
}
