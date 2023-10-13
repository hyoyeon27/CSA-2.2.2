package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	//declare a new reader for conn
	msg, _ := reader.ReadString('\n')
	//read a single string
	fmt.Println(msg)
	//fmt.Println(*conn, "Server OK")
	//_ is returning the error
}

func main() {
	ln, _ := net.Listen("tcp", ":8030")
	//call net.Listen on a particular port via particular protocol
	for {
		conn, _ := ln.Accept()
		//accept the connection and assign it to conn
		go handleConnection(conn)
		//this is no longer a block function
		//can have multiple connections concurrently
	}
}
