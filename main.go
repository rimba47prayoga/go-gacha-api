package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimba47prayoga/go-gacha-api/models"
	"github.com/rimba47prayoga/go-gacha-api/views"
)


func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/character", views.CreateCharacter)
	apiV1.GET("/character/:id", views.GetCharacter)
	apiV1.PUT("/character/:id", views.UpdateCharacter)
	apiV1.DELETE("/character/:id", views.DeleteCharacter)

	apiV1.POST("/weapon", views.CreateWeapon)
	apiV1.POST("/wish", views.Wish)

	models.ConnectDatabase()

	r.Run()
}
