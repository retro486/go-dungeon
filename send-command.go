/*
Communication tool for sending commands to the director. Things like user
management and game management, i.e. regenerate the maze or somesuch. Currently
this tool does absolutely nothing; more of a place-holder.
*/
package main

import (
	"net"
	"flag"
	"json"
	"fmt"
	"os"
	"io"
	"strings"
	"regexp"
)

var (
	connect_to = flag.String("c", "localhost:7000", "connect to server:port")
)

func validateUserName(username string) (bool,string) {
	// just good enough to sanitize usernames
	name_regex := regexp.MustCompile("[a-z]+[a-z0-9]*")
	matched := name_regex.MatchStrings(strings.ToLower(username))
	if len(matched) == 0 {
		msg := username + ": invalid name"
		return false,msg
	}
	return true,matched[0]
}

func processUserVerb(conn io.ReadWriteCloser, args []string, verb int) string {
	is_valid,username := validateUserName(args[1])
	if !is_valid { return username }
	
	uverb := new(UserVerb)
	uverb.Verb = verb
	uverb.Name = username
	
	// do additional testing of verb type and special handling here
	
	json_data,err := json.Marshal(&uverb)
	checkErr("Unable to encode data:", err)
	
	return sendJSON(conn, json_data, UVERB)
}

func parseArgs(args []string) string {
	conn,err := net.Dial("tcp", "", *connect_to)
	checkErr("Unable to connect:", err)
	
	response := ""
	switch strings.ToLower(args[0]) {
	default: response = "Invalid command."; break;
	case "server-shutdown":
		response = "Not yet implemented."
		break;
	case "server-restart":
		response = "Not yet implemented."
		break;
	case "user-create":
		if len(args) < 5 { return "MISSING PARAMETERS" }
		response = processUserVerb(conn, args, UVERB_CREATE)
		break;
	case "user-delete":
		if len(args) < 2 { return "MISSING USER NAME" }
		response = processUserVerb(conn, args, UVERB_DELETE)
		break;
	}
	
	conn.Close()
	return response
}

func main() {
	if flag.NArg() == 0 {
		checkErr("", os.NewError("No commands given."))
	} else if flag.NArg() > 5 {
		checkErr("", os.NewError("Too many commands given."))
	}

	fmt.Println(parseArgs(flag.Args()))
}

