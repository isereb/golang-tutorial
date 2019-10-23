package main

import (
	"fmt"
	"log"
	"net"
)

var port = 8081
var maxConnections = 3

func main() {

	log.Println("Starting a program at port:", port)
	log.Println("Program will end when it will reach", maxConnections, "connections")

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

	counter := 0
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic("Failed to accept connection due to an error: ", err)
		}

		_, err = conn.Write([]byte("Kiss my metal shiny ass\n"))
		if err != nil {
			log.Panic("Failed to write to connection due to an error: ", err)
		}

		err = conn.Close()
		if err != nil {
			log.Panic("Failed to close the connection due to an error: ", err)
		}

		counter++

		log.Println("Total amount of connections received:", counter)

		// After reaching max number of connections close the listener
		if counter == maxConnections {
			log.Println("Reached maximum amount of connections (", maxConnections, "). Finishing the execution")
			break
		}

	}

}
