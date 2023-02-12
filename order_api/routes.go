package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAllRoutes(db *gorm.DB, r *gin.RouterGroup) {
	AddRestaurantUserRoutes(db, r)
	AddMenuRoutes(db, r)
}
