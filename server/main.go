package main

import (
	"server/db"
	"server/models"
	"server/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.Post{})
	db.DB.AutoMigrate(&models.Vote{})

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))
	api := e.Group("/api")

	routers.RegisterPostRoutes(api.Group("/posts"))

	e.Logger.Fatal(e.Start(":1234"))
}
