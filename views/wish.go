package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimba47prayoga/go-gacha-api/utils"
)


type WishSerializer struct {
	Value	uint8	`json:"value" binding:"required"`
}


func Wish(ctx *gin.Context) {
	var data WishSerializer
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var response []interface{}
	for i := 0; i < int(data.Value); i++ {
		response = append(response, utils.Gacha())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"results": response,
	})
}
