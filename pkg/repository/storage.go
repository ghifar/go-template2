package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gotdd/pkg/api"
	"log"
	"path/filepath"
	"runtime"
)

type Storage interface {
	RunMigrations(connString string) error
	CreateUser(req api.NewUserRequest) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) RunMigrations(connString string) error {
	if connString == "" {
		return errors.New("repository: the connString was empty")
	}

	//get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")
	log.Printf("base path: %s", basePath)

	migrationsPath := filepath.Join("file://", basePath, "/pkg/repository/migrations/")

	log.Printf("Migration path: %s", migrationsPath)
	log.Printf("connection string: %s", connString)
	m, err := migrate.New(migrationsPath, connString)

	if err != nil {
		return err
	}

	err = m.Up()
	switch err {
	case errors.New("no change"):
		return nil
	}

	return nil
}

func (s *storage) CreateUser(req api.NewUserRequest) error {
	newUserStatement :=
		`INSERT INTO "user" (name, age, height, sex, activity_level, email, weight_goal)
			VALUES ($1, $2, $3, $4, $5, $6, $7);`

	err := s.db.QueryRow(
		newUserStatement,
		req.Name,
		req.Age,
		req.Height,
		req.Sex,
		req.ActivityLevel,
		req.Email,
		req.WeightGoal,
	).Err()
	if err != nil {
		log.Printf("error create user storage %v", err)
		return fmt.Errorf("error create user storage %v", err)
	}
	return nil
}
