package main

import (
	"strings"
)

// scrambles the given string array of doors; if skip_first is set, it won't scramble the
// first door (i.e., to ensure the first door is always "go to previous room")
func mixMenuItems(menu []string, skip_first bool) []string {
	max_len := len(menu)

	for i := 0; i < max_len; i++ {
		rand_index := int(randomNumberBetween(uint32(i), uint32(max_len)))
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

func mazeListDoors(node *mazeNode) []string {
	door_list := new([MAZE_MAX_DOORS]string)
	
	index := 0
	is_entrance := false
	
	// unavoidable special case for the entrance node
	if strings.ToUpper(node.name) != "ENTRANCE" {
		door_list[index] = "The previous room";
		index++
	} else { is_entrance = true }
	
	if len(node.doors) > 1 {
		for index < len(node.doors) {
			door_list[index] = node.doors[index].name
			index++
		}
	}
	
	return mixMenuItems(door_list[0:index], !is_entrance) // DON'T skip first item in mixup if entrance.
}

func mazeEnterDoor(node *mazeNode, door_num int) ([]string,*mazeNode) {
	if door_num >= len(node.doors) || door_num < 0 {
		return []string{"ERROR:BAD DOOR NUMBER"}, node
	}
	
	node = node.doors[door_num]
	if node.event_callback != nil && node.event_ran == false {
		message,_ := node.event_callback()
		node.event_ran = true
		// TODO do something with game over
		return []string{message}, node
	}
	return []string{"OK"},node
}

func loot(node *mazeNode, backpack []InventoryItem) ([]string,*mazeNode,
	[]InventoryItem) {
	// TODO assumes there's nothing to loot by default:
	// nothing was lootable, add some dirt to player's inventory
	return []string{"You looted some dirt because that's all what was available, you " + 
		"clepto."},node,backpack
}

