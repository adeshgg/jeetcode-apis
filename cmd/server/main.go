package main

import (
	"fmt"
	"jeetcode-apis/config"
	"jeetcode-apis/internal/cache"
	"jeetcode-apis/internal/db"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/local.config.json")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	cache, err := cache.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the redis: %v", err)
	}

	fmt.Println(database)
	fmt.Println(cache)
	// send database & cache to the problem service

	// Todo
	// router := api.SetupRouter(todoService, postService)

	// log.Println("Server running on port 8080")
	// router.Run(":8080")
}
