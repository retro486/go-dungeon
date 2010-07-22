package main

import (
	"fmt"
)

/*
Event functions that are to be called from mazeNodes; _MUST_ return a boolean
value representing whether or not to end the game.
*/

func onEnter() bool {
	fmt.Println("Welcome to the dungeon!")
	return false
}

func onExit() bool {
	fmt.Println("You've escaped the dungeon!")
	return true
}

