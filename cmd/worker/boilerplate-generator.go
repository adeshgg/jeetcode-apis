package worker

import (
	"jeetcode-apis/internal/cache"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

func StartWorker(rdb *redis.Client, key string, numWorkers int) ([]chan bool, *sync.WaitGroup) {
	wg := &sync.WaitGroup{}
	doneChannels := make([]chan bool, numWorkers)

	for i := 0; i < numWorkers; i++ {
		done := make(chan bool)
		doneChannels[i] = done
		wg.Add(1)
		go worker(i+1, rdb, key, done, wg)
	}

	return doneChannels, wg
}

func worker(id int, rdb *redis.Client, key string, done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-done:
			log.Printf("Worker: %d is stopping", id)
			return

		default:
			task, err := cache.DequeueTask(rdb, key)
			if err != nil {
				log.Printf("Worker %d: Error consuming task: %v", id, err)
				continue
			}
			if task != "" {
				log.Printf("Worker %d: processing task: %s", id, task)
				processTask(task)
			} else {
				log.Printf("Worker %d: no task found, retrying after 1 Second...", id)
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func processTask(task string) {
	log.Printf("Processing task: %s", task)
	// time.S
}

// Stop workers by signaling them to stop
func StopWorkers(doneChannels []chan bool) {
	for _, done := range doneChannels {
		done <- true // Send stop signal to workers
	}
}
