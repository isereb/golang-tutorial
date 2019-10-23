package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var port = 8081

func main() {
	log.Println("Starting a program at port:", port)

	// Start listening to port 8080
	li, err := net.Listen("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		log.Panic("Could not open a connection on port ", port, " because of error: ", err)
	}

	// Close the listener at the end of the execution
	defer func() {
		err = li.Close()
		if err != nil {
			log.Panic("Failed to close connection listener due to an error: ", err)
		}
		log.Println("Done.")
	}()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic("Failed to accept connection due to an error: ", err)
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		_, err := fmt.Fprintf(conn, "Recieved message: %s\n", ln)
		if err != nil {
			log.Panic("Failed to write to connection due to an error: ", err)
		}
	}
}
