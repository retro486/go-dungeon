package main

import (
	"flag"
	"net"
	"fmt"
	"os"
)

var (
	safe_mode = flag.Bool("s", false, "run in safe mode")
	bind_to = flag.String("b", "localhost:7000",
		"ip:port to bind to. can use hostname as well.")
	rand_enc_on_roll = flag.Bool("r", false,
		"random encounter chances are generated on each roll")
)

func main() {
	var conn net.Conn
	var response string
	var eot bool
	var node *mazeNode
	var player *Player
	
	flag.Parse()

	if *safe_mode && (*bind_to != "localhost:7000" || *rand_enc_on_roll) {
		checkErr("", os.NewError("Can't specify safe mode AND set a binding " +
			" or specify random encounters."))
	}
	
	fmt.Println("Starting up...")
	
	fmt.Print("Generating dungeon")
	node = mazeGenerateExitPath()
	node.name = "Entrance"
	mazeGenerateExtraPaths(node, 0)
	fmt.Println(".") // terminate progress indicator
	
	if *rand_enc_on_roll {
		fmt.Println("Random encounters to be generated on each roll; " +
			"skipping")
	} else {
		fmt.Print("Generating random encounters")
		mazeGenerateRandomEncounters(node)
	}
	fmt.Println(".") // terminate progress indicator

	fmt.Println("Binding to " + *bind_to + "...")
	server,err := net.Listen("tcp", *bind_to)
	checkErr("Unable to bind:", err)
	defer server.Close()
	
	// assemble a player
	player = new(Player)
	player.inventory = make([]InventoryItem, 20) // max 20 items
	player.weapons = make([]Weapon, 2) // max two weapons
	player.armor = make([]Armor, 5) // head, feet, legs, chest, hands
	player.position = node
	
	conn, err = server.Accept()
	checkErr("Problem accepting:", err)
	fmt.Println("Ready!")
	for {
		response,eot,node = doReading(conn,player)
		// if transmission has not yet ended
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

