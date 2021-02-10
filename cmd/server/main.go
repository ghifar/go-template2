package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gotdd/pkg/api"
	"gotdd/pkg/app"
	"gotdd/pkg/repository"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := "postgres://postgres:root@localhost/gotdd?sslmode=disable"
	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	//storage dependency
	storage := repository.NewStorage(db)

	//run migrations script
	err = storage.RunMigrations(connectionString)
	if err != nil {
		return err
	}

	//user service
	userService := api.NewUserService(storage)

	//weight service
	weightService := api.NewWeightService(storage)

	//router dependency
	router := gin.Default()
	router.Use(cors.Default())

	//start server
	server := app.NewServer(router, userService, weightService)
	err = server.Run()

	if err != nil {
		return err
	}
	return nil
}

func setupDatabase(connectionString string) (*sql.DB, error) {
	//initialize db driver
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	//ping to ensure it's connected
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
