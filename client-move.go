package main

import (
	"net"
	"flag"
	"json"
	"fmt"
	"os"
)

var (
	connect_to = flag.String("c", "localhost:7000", "connect to server:port")
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		checkErr("",os.NewError("No door number specified."))
	}	

	conn,err := net.Dial("tcp", "", *connect_to)
	checkErr("Unable to connect:", err)
	
	// the actual command to send:
	verb_move := new(PlayerVerb)
	verb_move.Verb = PVERB_MOVE
	verb_move.Parameter = flag.Args()[0]

	json_data,err := json.Marshal(&verb_move)
	checkErr("Unable to encode data:", err)
	
	fmt.Println(sendJSON(conn, json_data, PVERB))

	conn.Close()	
}

