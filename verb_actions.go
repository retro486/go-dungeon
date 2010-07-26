package main

func mazeListDoors(node *mazeNode) []string {
	door_list := new([MAZE_MAX_DOORS]string)
	index := 0
	if node.parent != nil {
		door_list[index] = "The previous room"; index++
	}
	for i := 0; i < len(node.doors); i++ {
		door_list[index] = node.doors[i].name; index++
	}

	// TODO scramble the door order
	
	return door_list
}
/*
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
*/
