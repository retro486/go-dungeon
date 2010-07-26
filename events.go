package main

/*
Event functions that are to be called from mazeNodes; _MUST_ return a boolean
value representing whether or not to end the game.
*/

func onEnter() (string,bool) {
	return "Welcome to the dungeon!",false
}

func onExit() (string,bool) {
	return "You've escaped the dungeon!",true
}

