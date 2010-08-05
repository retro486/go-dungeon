package main

import (
	"json"
	"io"
	"strconv"
)

// Various general verbs; be sure to test what verb type is stored as not all
// the parameters are used in all verbs!
type PlayerVerb struct {
	Verb int
	Parameter string
}
type UserVerb struct {
	Verb,Agility,Intelligence,Strength int
	Name string
}
type ServerVerb struct {
	Verb int
	Parameter string
}

const (
	// verb types
	PVERB = iota
	SVERB
	UVERB
	
	// player verbs
	PVERB_LOOK 
	PVERB_MOVE
	PVERB_ATTACK
	PVERB_EQUIP
	PVERB_UNEQUIP
	PVERB_DROP
	PVERB_LOOT
	
	// server command verbs
	SVERB_SHUTDOWN
	SVERB_RESET
	
	// user command verbs
	UVERB_CREATE
	UVERB_DELETE
)

func handleServerVerb(sverb *ServerVerb) (string,bool) {
	switch sverb.Verb {
	case SVERB_SHUTDOWN:
		return "OK:SHUTTINGDOWN",true
	case SVERB_RESET:
		return "OK:RESET",true
	}
	return "FAIL:UNKNOWNSERVERVERB",false
}

func handleUserVerb(uverb *UserVerb) string {
	switch uverb.Verb {
	case UVERB_CREATE:
		return "OK:CREATEUSER:Not implemented"
	case UVERB_DELETE:
		return "OK:DELETEUSER:"+uverb.Name
	}
	return "FAIL:UNKNOWNUSERVERB"
}

func handlePlayerVerb(pverb *PlayerVerb, node *mazeNode) ([]string,*mazeNode) {
	switch pverb.Verb {
	case PVERB_LOOK:
		return mazeListDoors(node), node

	case PVERB_MOVE:
		door_num,err := strconv.Atoi(pverb.Parameter)
		checkErr("Unable to parse door number", err) //TODO this is recoverable
		return mazeEnterDoor(node, door_num)
	
	case PVERB_EQUIP:
	case PVERB_UNEQUIP:
	case PVERB_DROP:
	case PVERB_ATTACK:
	case PVERB_LOOT:
		str,node,backpack = loot(node,backpack)
		return str,node
	
	default:
		return []string{"OK:PLAYERVERBS:Not implemented"}, node
	}
	return []string{"FAIL:UNKNOWNPLAYERVERB"}, node
}

// Gets all the appropriate data from the client and returns a status message
func doReading(conn io.ReadWriteCloser, node *mazeNode) (string,bool,*mazeNode) {
	// by definition, we expect a simple header defined as:
	header := make([]byte, 2) // [0] = verb type, [1] = ENCODED verb size
	_,err := conn.Read(header)
	checkErr("Unable to read header:", err)

	// create the buffer
	raw_json := make([]byte, header[1])
	
	// now get the data for whatever is coming next
	num_bytes,err := conn.Read(raw_json)
	checkErr("Unable to read data:", err)
	
	// if there was no more data to accept, we're done
	if num_bytes == 0 { return "",true,node }
	
	switch header[0] {
	case PVERB:
		pverb := new(PlayerVerb)
		err := json.Unmarshal(raw_json, &pverb)
		checkErr("Unable to decode PlayerVerb:", err)
		response,node := handlePlayerVerb(pverb, node)
		response_json,err := json.Marshal(&response)
		checkErr("Unable to encode response:", err)
		return string(response_json),false,node
		
	case SVERB:
		sverb := new(ServerVerb)
		err := json.Unmarshal(raw_json, &sverb)
		checkErr("Unable to decode ServerVerb:", err)
		response,flag := handleServerVerb(sverb)
		return response,flag,node
		
	case UVERB:
		uverb := new(UserVerb)
			err := json.Unmarshal(raw_json, &uverb)
		checkErr("Unable to decode ServerVerb:", err)
		return handleUserVerb(uverb),false,node
	}
	return "FAIL:UNKNOWNVERB",false,node
	
}

