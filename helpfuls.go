package main

import (
	"os"
	"io"
)

func checkErr(message string, err os.Error) {
	// ignore os.EOF since it's not really an error...
	if err != nil && err != os.EOF {
		os.Stderr.WriteString(message + "\n" + err.String() + "\n")
		os.Exit(1)
	}
}

func sendJSON(conn io.ReadWriteCloser, json_data []byte, verb_type uint8) string {
	// send the two header bytes
	header := []byte{verb_type, uint8(len(json_data))} // len(json_data) MUST < 128
	_,err := conn.Write(header)
	checkErr("Unable to send header:", err)
	
	// send the actual json-encoded data
	_,err = conn.Write(json_data)
	checkErr("Unable to send data:", err)
	
	// wait for a response from the server
	response := make([]byte, 255)
	_,err = conn.Read(response)
	checkErr("Unable to read response:", err)
	
	return string(response)
}
