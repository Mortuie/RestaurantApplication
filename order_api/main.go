package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := NewDbClient("order.db")
	if err != nil {
		panic("Failed to connect to sqlitedb")
	}

	db.AutoMigrate(&RestaurantUser{})
	db.AutoMigrate(&Menu{})
	db.AutoMigrate(&LineItem{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&MenuSections{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routerV1 := r.Group("v1")
	AddAllRoutes(db, routerV1)

	fmt.Printf("API running on port: %d", 8000)
	r.Run(":8000")
}
