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
		node := mazeGenerateExitPath()
		mazeGenerateExtraPaths(node, 0)
		
		if *rand_enc_on_roll {
			fmt.Println("Random encounters to be generated on each roll; " +
				"skipping...")
		} else {
			fmt.Println("Generating random encounters...")
			//TODO
		}
		
		var conn net.Conn
		fmt.Println("Ready!")
		
		// start accepting
		var response string
		var eot bool
		
		conn, err = server.Accept()
		checkErr("Problem accepting:", err)
		for {
			response,eot,node = doReading(conn,node)
			// if transmission has not yet ended
			_ = node // stop go from stupidly complaining that node is never used.
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

