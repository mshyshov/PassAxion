package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mshyshov/PassAxion/pkg/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	log.Println("#################### INITIALIZATION START ####################")
	// Connect to the DB
	log.Println("Connecting to database")
	db, err := gorm.Open(sqlite.Open("passaxion.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}
	log.Println("Successfully connected to database")

	// Migrate current models
	err = entities.MigrateEntities(db)
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate entities: %v", err))
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/auth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "login successful"})
	})

	log.Println("##################### INITIALIZATION END #####################")
	log.Println("Starting web server")
	router.Run("0.0.0.0:9000")
}
