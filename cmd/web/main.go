package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mshyshov/PassAxion/pkg/config"
	"github.com/mshyshov/PassAxion/pkg/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Main"})
}

func ListApps(ctx *gin.Context) {
	var apps []entities.Application
	err := entities.GetAllApplications(&apps)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Applications", "appsList": apps})
}

func ListUsers(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Users"})
}

func CreateAppForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "app_form.tmpl", gin.H{"title": "Create Application", "action": "new", "app_name_value": "", "app_description_value": "", "app_uri_value": ""})
}

func CreateAppPost(ctx *gin.Context) {
	var app entities.Application
	app.Name = ctx.PostForm("name")
	app.Description = ctx.PostForm("description")
	app.URI = ctx.PostForm("uri")

	entities.CreateApplication(&app)
	ctx.Redirect(http.StatusFound, "/apps")
}

func main() {
	db, err := gorm.Open(sqlite.Open("passaxion.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}
	config.DB = db
	entities.MigrateEntities(config.DB)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", Index)
	router.GET("/apps", ListApps)
	router.GET("/users", ListUsers)
	router.GET("/apps/new", CreateAppForm)
	router.POST("/apps/new", CreateAppPost)

	router.Run("0.0.0.0:9000")
}
