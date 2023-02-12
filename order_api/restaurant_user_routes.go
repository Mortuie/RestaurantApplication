package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RURepository struct {
	db *gorm.DB
}

func (rur *RURepository) createRU(c *gin.Context) {
	var ru RestaurantUser
	var err error

	if err = c.ShouldBindJSON(&ru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if err = ru.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if rur.db.Create(&ru).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Creation failed"})
		return
	}

	c.JSON(http.StatusOK, &ru)
}

func AddRestaurantUserRoutes(db *gorm.DB, router *gin.RouterGroup) {
	ruRepository := RURepository{db: db}
	router.POST("/restaurant-user", ruRepository.createRU)
}
