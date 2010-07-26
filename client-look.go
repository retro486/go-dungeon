package main

import (
	"net"
	"flag"
	"json"
	"fmt"
)

var (
	connect_to = flag.String("c", "localhost:7000", "connect to server:port")
)

func main() {
	flag.Parse()

	conn,err := net.Dial("tcp", "", *connect_to)
	checkErr("Unable to connect:", err)
	
	// the actual command to send:
	verb_look := new(PlayerVerb)
	verb_look.Verb = PVERB_LOOK

	json_data,err := json.Marshal(&verb_look)
	checkErr("Unable to encode data:", err)
	
	fmt.Println(sendJSON(conn, json_data, PVERB))

	conn.Close()	
}

