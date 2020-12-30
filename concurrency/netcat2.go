package concurrency


// Netcat1 is a read-only TCP client. package main
import ( "io"
	"log"
	"net"
	"os"
)
func NetCat2() {
	conn, err := net.Dial("tcp", "localhost:7111")
	if err != nil {
	}
	log.Fatal(err)
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
