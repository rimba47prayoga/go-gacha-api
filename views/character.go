package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimba47prayoga/go-gacha-api/models"
)


type CharacterSerializer struct {
	Name	string	`json:"name" binding:"required"`
	Rarity	uint8	`json:"rarity" binding:"required"`
	Element	string	`json:"element" binding:"required"`
	WeaponType	string	`json:"weapon_type" binding:"required"`
}


func CreateCharacter(ctx *gin.Context) {
	var serializer CharacterSerializer
	if err := ctx.ShouldBindJSON(&serializer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	char := models.Character{
		BaseItem: models.BaseItem{
			Name: serializer.Name,
			Rarity: serializer.Rarity,
		},
		Element: serializer.Element,
		WeaponType: serializer.WeaponType,
	}
	models.DB.Create(&char)
	ctx.JSON(http.StatusCreated, char)
}

func UpdateCharacter(ctx *gin.Context) {
	var serializer CharacterSerializer
	if err := ctx.ShouldBindJSON(&serializer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	char_id := ctx.Param("id")
	var char models.Character
	if query := models.DB.Find(&char, char_id); query.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	char.BaseItem.Name = serializer.Name
	char.BaseItem.Rarity = serializer.Rarity
	char.Element = serializer.Element
	char.WeaponType = serializer.WeaponType
	models.DB.Save(&char)
	ctx.JSON(http.StatusOK, char)
}

func GetCharacter(ctx *gin.Context) {
	char_id := ctx.Param("id")
	var char models.Character
	if query := models.DB.Find(&char, char_id); query.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	ctx.JSON(http.StatusOK, char)
}

func DeleteCharacter(ctx *gin.Context) {
	char_id := ctx.Param("id")
	var char models.Character
	if query := models.DB.Find(&char, char_id); query.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	models.DB.Delete(&char)
	ctx.JSON(http.StatusNoContent, gin.H{"success": true})
}
