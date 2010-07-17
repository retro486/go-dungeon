package main

import (
	"flag"
	"net"
	"fmt"
)

var (
	safe_mode = flag.Bool("s", false, "run in safe mode")
	bind_to = flag.String("b", "localhost:7000",
		"ip:port to bind to. can use hostname as well.")
	rand_enc_on_roll = flag.Bool("r", false,
		"random encounter chances are generated on each roll")
)

func main() {
	flag.Parse()
	
	if *safe_mode {
		// safe mode init
		fmt.Print("Starting up in safe mode...\n")
		fmt.Print("Binding to safe localhost:7000...\n")
		fmt.Print("Generating random encounters...\n")
	} else {
		// normal init
		fmt.Print("Starting up...\n")
		fmt.Print("Binding to " + *bind_to + "...\n")
		
		server,err := net.Listen("tcp", *bind_to)
		checkErr("Unable to bind:", err)
		defer server.Close()
		
		if *rand_enc_on_roll {
			fmt.Print("Random encounters to be generated on each roll; " +
				"skipping...\n")
		} else {
			fmt.Print("Generating random encounters...\n")
		}
		
		var conn net.Conn
		fmt.Print("Listening for connections on " + *bind_to + "...\n")
		
		// start accepting
		conn, err = server.Accept()
		checkErr("Problem accepting:", err)
		for {
			response,eot := doReading(conn)
			// if transmission has not yet ended
			if !eot {
				fmt.Println(response)
				// send response to client
				conn.Write([]byte(response))
			} else {
				// restart the connection; the client has disconnected
				conn.Close()
				conn,err = server.Accept()
				checkErr("Problem accepting:", err)
			}
		}
	}
}

