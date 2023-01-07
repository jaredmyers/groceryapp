package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	routes "github.com/jaredmyers/groceryapp/product_service/consumer/routes"
	"github.com/segmentio/kafka-go"
)

type info struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

var testInfo = []info{
	{ID: 1, Item: "Raisin Bagel"},
	{ID: 2, Item: "Milk"},
	{ID: 3, Item: "Waffles"},
	{ID: 4, Item: "Syrup"},
	{ID: 5, Item: "Mexican Cheese"},
	{ID: 6, Item: "Chicken Breast"},
	{ID: 7, Item: "Wrap"},
	{ID: 8, Item: "Hummus"},
	{ID: 9, Item: "Olive Oil"},
	{ID: 10, Item: "Bread"},
}

func main() {

	readFromKafka()

}

func readFromKafka() {
	conf := kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "products",
		Partition: 0,
	}
	r := kafka.NewReader(conf)
	//r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("%T\n", m.Value)
		fmt.Printf("message at offset %d: %s\n", m.Offset, string(m.Key), string(m.Value))
		writeToKafka("test", getData())
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

}

func getData() []byte {

	ctx, client := routes.DatabaseInstance()
	defer client.Disconnect(ctx)
	ps := routes.NewProductService(ctx, client.Database("product_service").Collection("products"))

	products, err := ps.GetProducts()
	if err != nil {
		log.Println(err)
	}

	log.Println("THIS IS PRODUCTS")
	log.Println(products)
	log.Printf("%T", products)
	for item := range products {
		log.Printf("%T, %v", item, item)
	}

	//ttt := &models.GroceryListPage{}

	jtestProducts, err := json.Marshal(products)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jtestProducts))

	//jtestInfo, err := json.Marshal(testInfo)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(jtestInfo))

	//return jtestInfo
	return jtestProducts
}

func writeToKafka(topic string, content []byte) {

	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
	fmt.Println("About to write: ")
	fmt.Println(string(content))

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("DBproducts"),
			Value: content,
		},
	)
	if err != nil {
		log.Fatal("failed to write message", err)
	}
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
