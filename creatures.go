package main

type InventoryItem struct {
	name string
	weight int
}

type Creature struct {
	hp,base_armor,base_dmg,level,
	strength,agility,intelligence,
	id int
	inventory []InventoryItem
	name string
}

type Player struct {
	*Creature
	current_xp int
}

type Monster struct {
	*Creature
	xp_worth int
}