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
	conn,err := net.Dial("tcp4", "", *connect_to)
	checkErr("Unable to connect:", err)
	
	data := "some command"
	json_data,err := json.Marshal(data)
	
	fmt.Println("Sending JSON data:\n",string(json_data))
	
	checkErr("Unable to encode JSON data:", err)
	_,err = conn.Write(json_data)
	checkErr("Unable to write:", err)
	conn.Close()
}
