package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	msgP := flag.String("msg", "Default message", "The message you want to send")
	flag.Parse()
	conn, _ := net.Dial("tcp", "172.31.22.3:8030")
	fmt.Fprintln(conn, *msgP)
	//flag is a package that stores values at the pointer
	//should use * in front of the variable to access the value, otherwise address
}
