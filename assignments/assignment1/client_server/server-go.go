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
	portString := ":" + serverPort
	ln, err := net.Listen("tcp", portString) // 1. Listen func

	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("creating new handlelConnection")
		go handelConnection(conn)
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
	defer c.Close()
	buf := make([]byte, RECV_BUFFER_SIZE)

	for {
		n, err := c.Read(buf) // 1. Read func
		if n > 0 {
			fmt.Fprint(os.Stdout, string(buf[:n]))
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
}
