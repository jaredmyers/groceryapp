// Service Layer
// Gin Server Golang

package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type info struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

type addTest struct {
	Selections []int `json:"selections"`
}

var newAdd addTest

var testInfo = []info{
	{ID: 1, Item: "Raisin Bagel"},
	{ID: 2, Item: "Milk"},
	{ID: 3, Item: "Waffles"},
	{ID: 4, Item: "Syrup"},
	{ID: 5, Item: "Mexican Cheese"},
	{ID: 6, Item: "Chicken Breast"},
	{ID: 7, Item: "Wrap"},
	{ID: 8, Item: "Hummus"},
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.GET("/test", getTestInfo)
	router.POST("/add", addTestInfo)
	router.GET("/testadd", getTestAdd)
	router.Run("localhost:8000")
}

func getTestAdd(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, newAdd)
}
func getTestInfo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, testInfo)
}
func addTestInfo(context *gin.Context) {

	if err := context.BindJSON(&newAdd); err != nil {
		log.Print(err)
		return
	}

	context.IndentedJSON(http.StatusCreated, newAdd)

}
