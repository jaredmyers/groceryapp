// Service Layer
// Gin Server Golang

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type groceryPageProduct struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

type addTest struct {
	Selections []int `json:"selections"`
}

var newAdd addTest

var testInfo = []groceryPageProduct{
	{ID: 1, Item: "Raisin Bagel"},
	{ID: 2, Item: "Milk"},
	{ID: 3, Item: "Waffles"},
	{ID: 4, Item: "Syrup"},
	{ID: 5, Item: "Mexican Cheese"},
	{ID: 6, Item: "Chicken Breast"},
	{ID: 7, Item: "Wrap"},
	{ID: 8, Item: "Hummus"},
}

func notmain() {
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

	triggerKafkaEvent("products")
	productsFromDb := readFromKafka("test")

	products := []groceryPageProduct{}

	err := json.Unmarshal(productsFromDb, &products)
	if err != nil {
		log.Panic(err)
	}
	context.IndentedJSON(http.StatusOK, products)
	//context.IndentedJSON(http.StatusOK, testInfo)
}
func addTestInfo(context *gin.Context) {

	if err := context.BindJSON(&newAdd); err != nil {
		log.Println(err)
		return
	}

	context.IndentedJSON(http.StatusCreated, newAdd)

}

func readFromKafka(topic string) []byte {
	conf := kafka.ReaderConfig{
		Brokers:     []string{"172.19.255.200:9092"},
		Topic:       topic,
		Partition:   0,
		StartOffset: kafka.LastOffset,
	}

	r := kafka.NewReader(conf)

	m, err := r.ReadMessage(context.Background())
	if err != nil {
		return []byte{}
	}

	log.Printf("%T\n", m.Value)
	log.Println(m.Value)

	return m.Value
}

func triggerKafkaEvent(topic string) {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("172.19.255.200:9092"),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("DBproducts"),
			Value: []byte("getDBproducts"),
		},
	)
	if err != nil {
		log.Fatal("failed to write message", err)
	}
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer", err)
	}

}
