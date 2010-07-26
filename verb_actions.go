package main

func mazeListDoors(node *mazeNode) []string {
	door_list := new([MAZE_MAX_DOORS]string)
	
	var has_parent bool
	index := 0
	
	// TODO: remove the special parent node; isn't neccessary.
	if node.parent != nil {
		door_list[index] = "The previous room"; index++; has_parent = true
	}
	for i := 0; i < len(node.doors); i++ {
		door_list[index] = node.doors[i].name; index++; has_parent = false
	}
	
	return mixMenuItems(door_list[0:index], has_parent)
}

