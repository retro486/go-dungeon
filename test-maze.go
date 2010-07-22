package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func printdoors(node *mazeNode) {
	if node.parent != nil {
		fmt.Println("0) The previous room")
	}
	for i := 0; i < len(node.doors); i++ {
		fmt.Printf("%d) %s\n", (i+1), node.doors[i].name)
	}
}
func look(node *mazeNode) {
	fmt.Println("You see:")
	printdoors(node)
}

func opendoor(node *mazeNode) *mazeNode {
	fmt.Print("Enter which room? ")
	door_num_str := getkbinput()
	door_num,err := strconv.Atoi(door_num_str)
	checkErr("Bad number entered:", err)
	// some simple error checking
	if door_num > len(node.doors) || (door_num == 0 && node.parent == nil) {
		fmt.Println("That door wasn't an option.")
		return node
	}
	if door_num == 0 { return node.parent }
	return node.doors[door_num-1]
}

func getkbinput() string {
	var buf []byte
	get_kb := bufio.NewReader(os.Stdin)
	buf,err := get_kb.ReadBytes('\n')
	checkErr("Unable to read STDIN:", err)
	command := strings.ToUpper(string(buf[0:len(buf)-1])) // trim trailing \n
	return command
}

func main() {
	var command string
	
	entrance := mazeGenerateExitPath()
	mazeGenerateExtraPaths(entrance, 0)
	
	node := entrance
	
	for done := false; !done; {
		if node.event_callback != nil && !node.event_ran {
			node.event_ran = true
			done = node.event_callback()
		}
		
		fmt.Print("Command: ")
		switch command = getkbinput(); command {
			case "EXIT": done = true; break
			case "LOOK": look(node); break
			case "GO": node = opendoor(node); break
			default: break
		}
	}
	
	fmt.Println("Exiting...")
}

