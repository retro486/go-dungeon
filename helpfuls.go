package main

import (
	"os"
)

func checkErr(message string, err os.Error) {
	if err != nil {
		os.Stderr.WriteString(message + "\n" + err.String() + "\n")
		os.Exit(1)
	}
}
