package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rimba47prayoga/go-gacha-api/models"
	"github.com/rimba47prayoga/go-gacha-api/views"
)


func main() {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/character", views.CreateCharacter)
	apiV1.GET("/character/:id", views.GetCharacter)
	apiV1.PUT("/character/:id", views.UpdateCharacter)
	apiV1.DELETE("/character/:id", views.DeleteCharacter)

	apiV1.POST("/weapon", views.CreateWeapon)
	apiV1.POST("/wish/weapon", views.WishWeapon)
	apiV1.GET("/wish/weapon/history", views.WishWeaponHistory)

	models.ConnectDatabase()

	r.Run()
}
