package main

import (
	"jeetcode-apis/api"
	"jeetcode-apis/config"
	"jeetcode-apis/internal/cache"
	"jeetcode-apis/internal/db"
	"jeetcode-apis/internal/service"
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

	problemService := service.NewProblemService(database, cache)
	router := api.SetupRouter(problemService)

	router.Run(":8080")
}
