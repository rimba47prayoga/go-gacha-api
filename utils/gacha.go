package utils

import (
	"math/rand"
	"time"

	"github.com/rimba47prayoga/go-gacha-api/models"
)


func GetRNGForRarity() StarRarity {
	// Create a new random source and generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomInt := rng.Intn(100)
	rarity := ThreeStar
	switch {
	case randomInt > int(FourStar):
		rarity = ThreeStar
	case randomInt <= int(FourStar) && randomInt > int(FiveStar):
		rarity = FourStar
	case randomInt <= int(FiveStar):
		rarity = FiveStar
	}
	return rarity
}


func GetPity() (uint8, uint8) {
	var four_star_pity uint8
	var five_star_pity uint8
	var last_history models.GachaHistory
	models.DB.Last(&last_history)
	if last_history.Rarity == uint8(FourStar) {
		four_star_pity = 0
	} else if last_history.Rarity == uint8(FiveStar) {
		five_star_pity = 0
	} else {
		var last_four_star models.GachaHistory
		result := models.DB.Where("rarity = ?", FourStar).Last(&last_four_star)

		// for empty result when first time gacha and didnt get any 4 star
		if result.Error != nil || result.RowsAffected == 0 {
			var count int64
			models.DB.Find(&models.GachaHistory{}).Count(&count)
			four_star_pity = uint8(count)
		} else {
			four_star_pity = uint8(last_history.ID - last_four_star.ID)
		}
		
		var last_five_star models.GachaHistory
		result = models.DB.Where("rarity = ?", FiveStar).Last(&last_five_star)
		// for empty result when first time gacha and didnt get any 5 star
		if result.Error != nil || result.RowsAffected == 0 {
			var count int64
			models.DB.Find(&models.GachaHistory{}).Count(&count)
			five_star_pity = uint8(count)
		} else {
			five_star_pity = uint8(last_history.ID - last_five_star.ID)
		}
	}
	return four_star_pity, five_star_pity
}


func GachaWeapon() (interface{}, StarRarity) {
	var rarity StarRarity
	four_star_pity, five_star_pity := GetPity()
	if (four_star_pity + 1) == uint8(FourStarGuarantee) {
		rarity = FourStar
	} else if (five_star_pity + 1) >= uint8(FiveStarGuarantee) {
		rarity = FiveStar
	} else {
		rarity = GetRNGForRarity()
	}
	var weapon models.Weapon
	models.DB.Where("rarity = ?", rarity).Order("RANDOM()").First(&weapon)
	models.DB.Create(&models.GachaHistory{
		WeaponID: &weapon.ID,
		Rarity: uint8(rarity),
	})
	return weapon, StarRarity(rarity)
}


func GachaCharacter() {}
