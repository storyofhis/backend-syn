package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/storyofhis/author-service/internal/repository/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(model.Author{})
	return db, nil
}
