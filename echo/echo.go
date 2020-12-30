package echo

import (
	"fmt"
	"os"
	"strings"
)

func EchoExample(){
	var s, sep string
	s = os.Args[0]
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep =""
	}
	fmt.Println(strings.Join(os.Args[1:], " "))
}
