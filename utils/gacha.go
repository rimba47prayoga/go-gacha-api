package utils

import (
	"math/rand"
	"time"

	"github.com/rimba47prayoga/go-gacha-api/models"
)

func Gacha() interface{} {
	// Create a new random source and generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomInt := rng.Intn(100)
	rarity := 60
	switch {
	case randomInt > 20:
		rarity = 60
	case randomInt <= 20 && randomInt > 1:
		rarity = 20
	case randomInt == 1:
		rarity = 1
	}
	var weapon models.Weapon
	models.DB.Where("rarity = ?", rarity).Order("RANDOM()").First(&weapon)
	return weapon
}
