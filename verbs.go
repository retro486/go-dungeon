package main

/*
Verbs (global vars to follow):
-Player
	look <direction>
	move <direction>
	attack <direction>
	loot
	drop <inventory id>
	equip <inventory id>
	unequip <inventory id>

-Director
	shutdown
	restart
	
-Account
	create <agility> <intelligence> <strength> <name>
	delete <name>
*/

import (
	"json"
	//"fmt"
)

type SingleVerb struct {
	Sverb,Parameter int
}

type PlayerVerb struct {
	Pverb,Agility,Intelligence,Strength int
	Name string
}

type JSONData struct {
	*SingleVerb
	*PlayerVerb
}

// TODO: do a better job at the embedded structs; shouldn't have to create copies of every
// single struct embedded... :(
func NewJSONData() *JSONData {
	return &JSONData{new(SingleVerb),new(PlayerVerb)}
}

func parseVerbs(json_data []byte) {
	// at this point we don't know what kind of struct this is
	//decoded_data := NewJSONData()
	
	// for now just use a plain string for receiving commands
	decoded_data := ""
	err := json.Unmarshal(json_data, &decoded_data)
	checkErr("Unable to parse JSON data:", err)
	
	/* the only way to tell what kind of struct this is, is to test presence of specific
		properties... */
	// TODO: do a better job at the embedded structs
	//fmt.Print("Parsed JSON data:\n",decoded_data.SingleVerb,"\n")
	
	// do something with the command
}
