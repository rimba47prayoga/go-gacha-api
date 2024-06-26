package models

import (
	"gorm.io/gorm"
)


type BaseItem struct {
	Name	string 	`db:"name"`
	Rarity	uint8 	`db:"rarity"`
}


type Character struct {
	gorm.Model
	BaseItem
	Element	string 	`db:"element"`
	WeaponType string `db:"weapon_type"`
}


type Weapon struct {
	gorm.Model
	BaseItem
	WeaponType	string `db:"weapon_type"`
}


type GachaHistory struct {
	gorm.Model
	CharacterID	*uint	`db:"character_id"`
	Character	*Character

	WeaponID	*uint	`db:"weapon_id"`
	Weapon		Weapon
	Rarity		uint8 	`db:"rarity"`  // for counting pity much easier
}
