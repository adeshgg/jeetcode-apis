package db

import (
	"fmt"
	// "todo-list-api/config"
	"jeetcode-apis/config"
	// "todo-list-api/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the database schemas
	// db.AutoMigrate(&model.Todo{}, &model.Post{})

	return db, nil
}
