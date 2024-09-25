package service

import (
	"jeetcode-apis/pkg/model"

	"github.com/go-redis/redis/v8"
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
	return p.db.Create(&problem).Error
}
