package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type mRepository struct {
	db *gorm.DB
}

func (mr *mRepository) createMenu(c *gin.Context) {
	var m Menu
	var err error

	if err = c.ShouldBindJSON(&m); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if err = m.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = mr.db.Create(&m).Error; err != nil {
		fmt.Println("error creating, ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Creation failed"})
		return
	}

	c.JSON(http.StatusOK, &m)
}

func (mr *mRepository) getSpecificMenu(c *gin.Context) {
	id := c.Params.ByName("id")
	menuId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	var menu Menu

	findErr := mr.db.First(&menu, menuId).Error

	if findErr != nil {
		fmt.Println(findErr, menu)
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func AddMenuRoutes(db *gorm.DB, router *gin.RouterGroup) {
	mRepository := mRepository{db: db}
	router.GET("/menu/:id", mRepository.getSpecificMenu)
	router.POST("/menu", mRepository.createMenu)

}
