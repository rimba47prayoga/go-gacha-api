package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimba47prayoga/go-gacha-api/models"
	"github.com/rimba47prayoga/go-gacha-api/utils"
)


type WishSerializer struct {
	Value	uint8	`json:"value" binding:"required"`
}


func WishWeapon(ctx *gin.Context) {
	var data WishSerializer
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var response []interface{}
	for i := 0; i < int(data.Value); i++ {
		result, _ := utils.GachaWeapon()
		response = append(response, &result)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"results": response,
	})
}

func WishWeaponHistory(ctx *gin.Context) {
	var response []models.GachaHistory
	pagination := utils.InitPagination(ctx)
	models.DB.Scopes(utils.Paginate(&response, pagination)).Preload("Weapon").Find(&response)
	pagination.Rows = response
	ctx.JSON(http.StatusOK, pagination)
}
