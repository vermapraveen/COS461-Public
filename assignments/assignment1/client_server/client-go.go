/*****************************************************************************
 * client-go.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"log"
	"net"
  "os"
  "io"
)

const SEND_BUFFER_SIZE = 2048

/* TODO: client()
 * Open socket and send message from stdin.
 */
func client(serverIP string, serverPort string) {
	conn, err := net.Dial("tcp", serverIP+":"+serverPort) // 1. Dial func

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, SEND_BUFFER_SIZE)

	for {
		n, err := os.Stdin.Read(buf) // 1. Stdin func
		if n > 0 {
			_, err := conn.Write(buf[:n]) // 1. Write func
			if err != nil {
				log.Fatalln(err)
      }
    }
    
    if err != nil{
      if err == io.EOF{
        break
      }

      log.Fatalln(err)
    }
	}

}

// Main parses command-line arguments and calls client function
func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ./client-go [server IP] [server port] < [message file]")
	}
	serverIP := os.Args[1]
	serverPort := os.Args[2]
	client(serverIP, serverPort)
}
