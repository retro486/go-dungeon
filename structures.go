package main

const (
	ARMOR_SLOT_HEAD = iota
	ARMOR_SLOT_SHOULDERS
	ARMOR_SLOT_CHEST
	ARMOR_SLOT_WRISTS
	ARMOR_SLOT_HANDS
	ARMOR_SLOT_LEGS
	ARMOR_SLOT_FEET
)

type InventoryItem struct {
	name string
	weight int
}

type Equippable struct {
	InventoryItem
	equipped bool
}

type Weapon struct {
	Equippable
	base_dmg int
	two_handed bool
}

type Armor struct {
	Equippable
	base_prot int
	slot int
}	

type Creature struct {
	hp,base_armor,base_dmg,level,
	strength,agility,intelligence,
	id int
	name string
	inventory []InventoryItem
	weapons []Weapon
	armor []Armor
	position *mazeNode
}

type Player struct {
	Creature // inherit general creature type information
	current_xp int
}

type Monster struct {
	Creature // inherit general creature type information
	xp_worth int
}

