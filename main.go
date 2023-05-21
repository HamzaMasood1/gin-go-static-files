package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name    string
	Picture string
}

func IndexHandler(c *gin.Context) {
	recipes := make([]Recipe, 0)
	recipes = append(recipes, Recipe{
		Name:    "Burger",
		Picture: "/assets/images/burger.jpg",
	})
	recipes = append(recipes, Recipe{
		Name:    "Pizza",
		Picture: "/assets/images/pizza.jpg",
	})
	recipes = append(recipes, Recipe{
		Name:    "Tacos",
		Picture: "/assets/images/tacos.jpg",
	})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"recipes": recipes,
	})
	// c.File("index.html")
}
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", IndexHandler)
	router.Run("localhost:8080")
}
