/*****************************************************************************
 * server-go.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const RECV_BUFFER_SIZE = 2048

/* TODO: server()
 * Open socket and wait for client to connect
 * Print received message to stdout
 */
func server(serverPort string) {
	portString := ":" + serverPort           // 1. go := vs =
	ln, err := net.Listen("tcp", portString) // 1. Listen func

	if err != nil {
		log.Fatal(err)
		return
	}

	for { // 1. for loop in go
		conn, err := ln.Accept() // 1. multi return in go
		if err != nil {          // 1. nil key world in go
			log.Fatal(err)
			return
		}

		go handelConnection(conn) // 1. go keyword
	}
}

// Main parses command-line arguments and calls server function
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./server-go [server port]")
	}
	serverPort := os.Args[1]
	server(serverPort)
}

func handelConnection(c net.Conn) {
	defer c.Close()                       // 1. defer in go 2. conn CLose func
	buf := make([]byte, RECV_BUFFER_SIZE) // 1. make func 2. func in go

	for {
		n, err := c.Read(buf) // 1. Read func
		if n > 0 {            // go if condition
			fmt.Fprint(os.Stdout, string(buf[:n])) // 1. fmt all print methods 2. string  method 3. array with :
		}

		if err != nil {
			if err == io.EOF { // 1. io namespace
				break
			}
			log.Fatal(err)
		}
	}
}
