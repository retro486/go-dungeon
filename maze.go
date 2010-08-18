/*
The maze generation is more complicated than simply four directions; nodes can
have up to nine direction choices. As such, none of the paths converge, even
if four left turns are made; each turn does not represent any standard distance
covered and the only to return to a previously explored area is to backtrack.
*/
package main

import (
	"rand"
	"time" // for Nanoseconds
	"fmt" // for printing progress indicators on the director side
)

const (
	MAZE_COLS = 10
	MAZE_ROWS = 10
	
	// chance modifiers; the higher the number, the bigger the chance.
	// Keep it below 100.
	MAZE_CHANCE_BATTLE = 85
	MAZE_CHANCE_LUCK = 10
	MAZE_CHANCE_TRAP = 50
	
	// for the randomizer for picking which event to put
	MAZE_CHANCE_TYPES = 3
	MCT_BATTLE = 0
	MCT_LUCK = 1
	MCT_TRAP = 2
)

type mazeCell struct {
	// caveat: can only have one kind of event per cell
	event_callback func() (string,bool)
	event_ran, gen_visited, is_wall bool
	north, south, east, west *mazeCell
}

func mazeGetCellAt(x, y int, maze *[]mazeCell) *mazeCell {
	pos := y + (x*y)
	return &maze[pos]
}

func mazeGenerate(maze *[]mazeCell) {
	var cell *mazeCell
	maze = make([]mazeCell, (MAZE_ROWS*MAZE_COLS))
	
	// initialize walls so they're not nil
	for i:=0; i<MAZE_ROWS; i++ {
		for j:=0; j<MAZE_COLS; j++ {
			cell = mazeGetCellAt(j, i, maze)
			if j == 0 {
				cell.north.is_wall = true
			}
			if i == 0 {
				cell.west.is_wall = true
			}
			if j == MAZE_COLS-1 {
				cell.east.is_wall = true
			}
			if i == MAZE_ROWS-1 {
				cell.south.is_wall = true
			}			
		}
	}
	
	current := &maze[0]
	curr_row := 0; curr_col := 0
	for visited := 1; visited < (MAZE_ROWS*MAZE_COLS); {
		fmt.Print(".")
		switch(randomNumber(4)) {
			case 0: if current.north == nil || !current.north.gen_visited {
				if current.north != nil && current.north.is_wall { continue }
				current.gen_visited = true
				current.north = mazeGetCellAt(curr_row-1, curr_col, maze)
				current = current.north
				curr_row--; visited++
				break
			} else {
				continue
			}
			case 1: if current.east == nil || !current.east.gen_visited {
				if current.east != nil && current.east.is_wall { continue }
				current.gen_visited = true
				current.east = mazeGetCellAt(curr_row, curr_col+1, maze)
				current = current.east
				curr_col++; visited++
				break;
			} else {
				continue
			}
			case 2: if current.south == nil || !current.south.gen_visited {
				if current.south != nil && current.south.is_wall { continue }
				current.gen_visited = true
				current.south = mazeGetCellAt(curr_row+1, curr_col, maze)
				current = current.south
				curr_row++; visited++
				break;
			} else {
				continue
			}
			case 3: if current.west == nil || !current.west.gen_visited {
				if current.west != nil && current.west.is_wall { continue }
				current.gen_visited = true
				current.west = mazeGetCellAt(curr_row-1, curr_col, maze)
				current = current.west
				curr_col--; visited++
				break;
			} else {
				continue
			}
		}
	}
	// maze now contains a valid maze with entry at (0,0) and some random exit
	return maze
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

// generates an event for the given maze cell
func mazeGenerateRandomEncounters(cell *mazeCell) {
	cell.event_ran = false // init the event_ran flag
	event_type := randomNumber(MAZE_CHANCE_TYPES)
	chance := randomNumber(100)
	switch event_type {
	case MCT_BATTLE:
		if chance <= MAZE_CHANCE_BATTLE { // arbitrarily pick a value to be "true"
			cell.event_callback = onBattle
		}
		break
	case MCT_LUCK:
		if chance <= MAZE_CHANCE_LUCK {
			cell.event_callback = onLuck
		}
		break
	case MCT_TRAP:
		if chance <= MAZE_CHANCE_TRAP {
			cell.event_callback = onTrap
		}
		break
	default: // no event -- honestly this should never be the case at this point.
		fmt.Println("Bad random: ", event_type) // DEBUG
		break
	}
	fmt.Print(".") // progress indicator
}

