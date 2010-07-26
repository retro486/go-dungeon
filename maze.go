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
)

const (
	//maybe randomly generated in the future to further randomize the maze
	MAZE_MAX_DOORS = 5
	MAZE_MAX_DEPTH = 9
	
	// maze event types
	MT_BATTLE = iota
	MT_TREASURE
	MT_TRAP
	MT_EXIT
	
	// face directions
	FD_NORTH
	FD_SOUTH
	FD_EAST
	FD_WEST
	FD_NONE // for marking an "empty" direction
)

type mazeNode struct {
	event_callback func() (string,bool)
	doors []*mazeNode
	parent *mazeNode // for more intuitive reference; no real reason it couldn't be doors[0].
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
	new_slice := make([]*mazeNode, (size+1))
	copy(new_slice, node.doors)
	node.doors = new_slice
	node.doors[size] = new(mazeNode)
	node.doors[size].parent = node
	return node.doors[size],nil
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
	// don't generate past MAZE_MAX_DEPTH levels
	if curr_level == MAZE_MAX_DEPTH { return }
	
	// add doors if there's room
	for i := len(enter.doors); i < MAZE_MAX_DOORS; i++ {
		if randomNumber(2) == 1 {
			new_door,err := createDoor(enter)
			checkErr("Unable to add a door:", err)
			new_door.name = mazeGenerateRandomName()
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

// scrambles the given string array of doors; if skip_first is set, it won't scramble the
// first door (i.e., to ensure the first door is always "go to previous room")
func mixMenuItems(menu []string, skip_first bool) []string {
	max_len := len(menu)

	for i := 0; i < max_len; i++ {
		rand_index := int(randomNumberBetween(uint32(i), uint32(max_len)))
		// TODO this isn't working or something...
		if skip_first && (i == 0 || rand_index == 0) {
			continue
		} else {
			temp := menu[i]
			menu[i] = menu[rand_index]
			menu[rand_index] = temp
		}
	}
	
	return menu
}

