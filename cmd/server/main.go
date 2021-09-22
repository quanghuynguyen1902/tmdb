package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tmdb/pkg/api"
	"github.com/tmdb/pkg/http"
	"github.com/tmdb/pkg/repository"

	"github.com/tmdb/pkg/db"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}

}

// func run will be responsible for setting up db connections, routers etc
func run() error {
	//// init database
	initDB, err := db.NewDatabase()
	if err != nil {
		return err
	}

	if err := initDB.SetupDBConnection(); err != nil {
		return err
	}

	defer initDB.CloseDatabaseConnection(initDB.DB)

	if err != nil {
		return err
	}

	//create storage dependency
	storage := repository.NewStorage(initDB.DB)

	cache, err := db.NewRedisCache(10)

	if err != nil {
		log.Fatalf("err init cache %v", err)
	}

	//// create router dependecy
	router := gin.Default()
	router.Use(cors.Default())

	//// create service
	service := api.NewService(storage, cache)

	//// create server
	server := http.NewServer(router, service)

	//// start the server
	err = server.Run()
	//
	if err != nil {
		return err
	}

	return nil
}
