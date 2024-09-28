package service

import (
	"encoding/json"
	"fmt"
	"jeetcode-apis/internal/cache"
	"jeetcode-apis/pkg/model"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProblemService struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProblemService(db *gorm.DB, redis *redis.Client) *ProblemService {
	return &ProblemService{
		db:    db,
		redis: redis,
	}
}

func (p *ProblemService) GetProblems() ([]model.Problem, error) {
	var problems []model.Problem

	if err := p.db.Find(&problems).Error; err != nil {
		return nil, err
	}

	return problems, nil
}

func (p *ProblemService) CreateProblems(problem model.Problem) error {

	// Start a new transaction
	tx := p.db.Begin()

	// Ensure that transaction is rolled back in case of a panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&problem).Error; err != nil {
		tx.Rollback()
		return err
	}

	// enqueue processing task to the redis queue
	problemTask := model.ProblemTask{
		ID:        uuid.New(),
		ProblemId: problem.ID,
		Link:      problem.Link,
	}

	// convert the task struct into a JSON string
	jsonTask, err := json.Marshal(problemTask)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error while marshaling struct to json: %v", err)
	}

	if err := cache.EnqueueTask(p.redis, "problem-queue", string(jsonTask)); err != nil {
		tx.Rollback()
		return fmt.Errorf("error while adding task to the queue: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error while committing transaction: %v", err)
	}

	return nil
}
