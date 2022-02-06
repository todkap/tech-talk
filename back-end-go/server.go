// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"log"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	cache "github.com/patrickmn/go-cache"
)

var localStorage = cache.New(5*time.Minute, 10 * time.Hour)

// MyStorage struct
type MyStorageItem struct {
	Key     string  `json:"k,omitempty" bson:"_id,omitempty"`
	Value   string  `json:"v"`
}
type MyLegacyStorageItem struct {
	Key     string  `json:"key,omitempty" bson:"_id,omitempty"`
	Value   string  `json:"value"`
}

type MyStorage struct {
	Items []MyStorageItem
}
func (storage *MyStorage) AddItem(item MyStorageItem) {
	storage.Items = append(storage.Items, item)
}


func main() {

	// Create a Fiber app
	app := fiber.New()
	
	// Get all items from storage service
	app.Get("/storage", func(c *fiber.Ctx) error {

		if(localStorage.ItemCount() == 0){
			return c.JSON(localStorage)
		}
		var cachedItems = localStorage.Items()
		fmt.Println("result size", len(cachedItems) )
		storage := MyStorage{}
		for key, element := range cachedItems {
			fmt.Println("Key:", key, "=>", "Element:", element.Object)
			storageItem := MyStorageItem { Key: key, Value: element.Object.(string) }
			storage.AddItem(storageItem)
		}
		return c.JSON(storage.Items)
	})

	// Get specific item from storage service
	app.Get("/storage/:id", func(c *fiber.Ctx) error {
		params := c.Params("id")

		val, found := localStorage.Get(params)
		if !found {
			return c.Status(500).SendString("Item not found")
		}
		return c.Status(fiber.StatusOK).JSON(val)
	})

	// add an item to storage
	app.Put("/storage", func(c *fiber.Ctx) error {
		storageItem := new(MyLegacyStorageItem)
		// Parse body into struct
		if err := c.BodyParser(storageItem); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		localStorage.Set(storageItem.Key, storageItem.Value, cache.DefaultExpiration)

		// return the created Employee in JSON format
		return c.Status(201).JSON(storageItem)
	})

	// Delete an item from storage
	app.Delete("/storage/:id", func(c *fiber.Ctx) error {
		localStorage.Delete(c.Params("id"))
		// the record was deleted
		return c.SendStatus(204)
	})
	log.Fatal(app.Listen(":9080"))
}