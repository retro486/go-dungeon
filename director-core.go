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
		fmt.Println("Starting up in safe mode...")
		fmt.Println("Binding to safe localhost:7000...")
		fmt.Println("Generating random encounters...")
	} else {
		// normal init
		fmt.Println("Starting up...")
		fmt.Println("Binding to " + *bind_to + "...")
		
		server,err := net.Listen("tcp", *bind_to)
		checkErr("Unable to bind:", err)
		defer server.Close()
		
		fmt.Println("Generating dungeon...")
		entrance := mazeGenerateExitPath()
		mazeGenerateExtraPaths(entrance, 0)
		
		if *rand_enc_on_roll {
			fmt.Println("Random encounters to be generated on each roll; " +
				"skipping...")
		} else {
			fmt.Println("Generating random encounters...")
		}
		
		var conn net.Conn
		fmt.Println("Ready!")
		
		// start accepting
		conn, err = server.Accept()
		checkErr("Problem accepting:", err)
		for {
			response,eot := doReading(conn,entrance)
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

