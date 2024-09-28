package main

import (
	"flag"
	"jeetcode-apis/api"
	"jeetcode-apis/cmd/worker"
	"jeetcode-apis/config"
	"jeetcode-apis/internal/cache"
	"jeetcode-apis/internal/db"
	"jeetcode-apis/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var workerType string

func main() {

	// Application flags
	flag.StringVar(&workerType, "worker-type", "api", "--worker-type=(api|boilerplate-generator)")
	flag.Parse()

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

	if workerType == "api" {
		router := api.SetupRouter(problemService)
		router.Run(":8080")
	} else if workerType == "boilerplate-generator" {
		doneChannels, wg := worker.StartWorker(cache, "problem-queue", 3)

		// Handle CTRL+C (SIGINT) signal for graceful shutdown
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sigs
			log.Printf("Received signal: %s. Stopping workers...", sig)
			worker.StopWorkers(doneChannels)
		}()

		wg.Wait()

		log.Println("All workers stopped. Exiting program.")
	}

}
