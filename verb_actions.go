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

func loot(node *mazeCell, backpack []InventoryItem) ([]string,*mazeCell,
	[]InventoryItem) {
	// TODO assumes there's nothing to loot by default:
	// nothing was lootable, add some dirt to player's inventory
	return []string{"You looted some dirt because that's all what was available, you " + 
		"clepto."},node,backpack
}

