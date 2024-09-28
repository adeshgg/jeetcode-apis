package model

import "github.com/google/uuid"

type ProblemCreateInput struct {
	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}

type ProblemTask struct {
	ID        uuid.UUID `json:"id"`
	ProblemId uint      `json:"problem_id"`
	Link      string    `json:"link"`
}
