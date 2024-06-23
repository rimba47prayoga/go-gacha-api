package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimba47prayoga/go-gacha-api/models"
)


type WeaponSerializer struct {
	Name	string	`json:"name" binding:"required"`
	Rarity	uint8	`json:"rarity" binding:"required"`
	WeaponType	string `json:"weapon_type" binding:"required"`
}


func CreateWeapon(ctx *gin.Context) {
	var serializer WeaponSerializer
	if err := ctx.ShouldBindJSON(&serializer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	weapon := models.Weapon{
		BaseItem: models.BaseItem{
			Name: serializer.Name,
			Rarity: serializer.Rarity,
		},
		WeaponType: serializer.WeaponType,
	}
	models.DB.Create(&weapon)
	ctx.JSON(http.StatusCreated, weapon)
}
