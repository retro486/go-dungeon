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

func onBattle() (string,bool) {
	return "You encountered a mob.",false
}

func onTrap() (string,bool) {
	return "You encountered a trap.",false
}

func onLuck() (string,bool) {
	return "You had good luck.",false
}

