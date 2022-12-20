// Service Layer
// Gin Server Golang

package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type info struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}

var testInfo = []info{
	{ID: "1", Item: "Raisin Bagel"},
	{ID: "2", Item: "Milk"},
	{ID: "3", Item: "Waffles"},
	{ID: "4", Item: "Syrup"},
	{ID: "5", Item: "Mexican Cheese"},
	{ID: "6", Item: "Chicken Breast"},
	{ID: "7", Item: "Wrap"},
	{ID: "8", Item: "Hummus"},
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.GET("/test", getTestInfo)
	router.Run("localhost:8000")
}

func getTestInfo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, testInfo)
}
