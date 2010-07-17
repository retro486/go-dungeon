package main

import (
	"os"
)

func checkErr(message string, err os.Error) {
	// ignore os.EOF since it's not really an error...
	if err != nil && err != os.EOF {
		os.Stderr.WriteString(message + "\n" + err.String() + "\n")
		os.Exit(1)
	}
}

