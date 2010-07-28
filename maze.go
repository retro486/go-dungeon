/*
The maze generation is more complicated than simply four directions; nodes can
have up to nine direction choices. As such, none of the paths converge, even
if four left turns are made; each turn does not represent any standard distance
covered and the only to return to a previously explored area is to backtrack.
*/
package main

import (
	"os"
	"rand"
	"time" // for Nanoseconds
	"fmt" // for printing progress indicators on the director side
)

const (
	//maybe randomly generated in the future to further randomize the maze
	MAZE_MAX_DOORS = 5
	MAZE_MAX_DEPTH = 9
	
	// chance modifiers; the higher the number, the bigger the chance. Keep it below 100.
	MAZE_CHANCE_BATTLE = 75 // 75% chance
	MAZE_CHANCE_LUCK = 10
	MAZE_CHANCE_TRAP = 40
	MAZE_CHANCE_TYPES = 3 // how many types of events are there (other than enter/exit)
	
	MCT_BATTLE = 0
	MCT_LUCK = 1
	MCT_TRAP = 2
	
	// face directions
	FD_NORTH = iota
	FD_SOUTH
	FD_EAST
	FD_WEST
	FD_NONE // for marking an "empty" direction
)

type mazeNode struct {
	event_callback func() (string,bool)
	doors []*mazeNode
	name string
	event_ran bool
}


func mazeGetColor() string {
	colors := []string {
		"green",
		"blue",
		"red",
		"brown",
		"black",
		"orange",
	}
	
	return colors[randomNumber(uint32(len(colors)))]
}

func mazeGetTexture() string {
	textures := []string {
		"crusty",
		"scaly",
		"rusty",
		"slimy",
		"ugly",
		"shiny",
		"decrepit",
		"new",
	}
	
	return textures[randomNumber(uint32(len(textures)))]
}

// appends a node to the given node and returns the new node
func createDoor(node *mazeNode) (*mazeNode,os.Error) {
	if node == nil {
		return nil,os.NewError("Given node is empty.")
	}
	
	size := len(node.doors)
	if size == MAZE_MAX_DOORS {
		return nil,os.NewError("Given node is full.")
	}
	
	// reallocate doors array and append a new node
	new_slice := make([]*mazeNode, (size+1), MAZE_MAX_DOORS)
	copy(new_slice, node.doors)
	node.doors = new_slice
	node.doors[size] = new(mazeNode)
	parent := node
	node = node.doors[size]
	node.doors = make([]*mazeNode, 1, MAZE_MAX_DOORS)
	node.doors[0] = parent
	return node,nil
}

func mazeGenerateRandomName() string {
	return (mazeGetTexture() + " " + mazeGetColor() + " door")
}

func mazeGenerateExitPath() *mazeNode {
	node := new(mazeNode)
	node.name = mazeGenerateRandomName()
	node.event_callback = onEnter
	
	var entrance *mazeNode = node
	var err os.Error = nil
	
	for i := 0; i < MAZE_MAX_DEPTH; i++ {
		node,err = createDoor(node)
		checkErr("Unable to add a door:", err)
		node.name = mazeGenerateRandomName()
	}
	
	node.event_callback = onExit
	
	return entrance
}

func mazeGenerateExtraPaths(enter *mazeNode, curr_level int) {
	if curr_level == MAZE_MAX_DEPTH	|| len(enter.doors) == MAZE_MAX_DOORS {
		return
	}
	
	// add doors if there's room
	for i := len(enter.doors); i < MAZE_MAX_DOORS; i++ {
		if randomNumber(2) == 1 { // 50% chance to add a door or not
			new_door,err := createDoor(enter)
			checkErr("Unable to add a door:", err)
			new_door.name = mazeGenerateRandomName()
			fmt.Print(".") // gives an idea of progress
		}
	}
	
	// enter each door and add doors from there
	for i := 0; i < len(enter.doors); i++ {
		mazeGenerateExtraPaths(enter.doors[i], curr_level+1)
	}
	
	return
}

// generates a random number between 0 and max_range-1
func randomNumber(max_range uint32) uint32 {
	source := rand.NewSource(time.Nanoseconds())
	rnd := rand.New(source)
	return (rnd.Uint32() % max_range)
}

func randomNumberBetween(low uint32, high uint32) uint32 {
	var num uint32
	for num = randomNumber(high); num < low && num > high; num = randomNumber(high){}
	return num
}

func mazeGenerateRandomEncounters(node *mazeNode) {
	if len(node.doors) == 1 { return } // node.doors will AT LEAST have one element for parent

	// start the index at [1] since [0] == parent and would infinite-loop this function.
	for i := 1; i < len(node.doors); i++ {
		door := node.doors[i]
		door.event_ran = false // init the event_ran flag
		event_type := randomNumber(MAZE_CHANCE_TYPES)
		chance := randomNumber(100)
		switch event_type {
		case MCT_BATTLE:
			if chance <= MAZE_CHANCE_BATTLE { // arbitrarily pick a value to be "true"
				door.event_callback = onBattle
			}
			break
		case MCT_LUCK:
			if chance <= MAZE_CHANCE_LUCK {
				door.event_callback = onLuck
			}
			break
		case MCT_TRAP:
			if chance <= MAZE_CHANCE_TRAP {
				door.event_callback = onTrap
			}
			break
		default: // no event -- honestly this should never be the case at this point.
			fmt.Println("Bad random: ", event_type) // DEBUG
			break
		}
		fmt.Print(".") // progress indicator
		mazeGenerateRandomEncounters(door) // recurse into this door
	}
}

