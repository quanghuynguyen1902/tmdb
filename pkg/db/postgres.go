package db

import (
	"fmt"
	"log"

	"github.com/tmdb/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {
	return &Database{}, nil
}

func (d *Database) SetupDBConnection() error {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("create new config error %s", err)
		return err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName)
	d.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to create a connection to database", err)
		return err
	}

	//d.DB.AutoMigrate(&models.Movie{})
	//d.DB.AutoMigrate(&models.Language{})
	//d.DB.AutoMigrate(&models.Country{})
	//d.DB.AutoMigrate(&models.Genre{})
	//d.DB.AutoMigrate(&models.Actor{})
	//d.DB.AutoMigrate(&models.Director{})
	//d.DB.AutoMigrate(&models.Person{})
	//d.DB.AutoMigrate(&models.Keyword{})
	//d.DB.AutoMigrate(&models.SpokenLanguage{})
	//d.DB.AutoMigrate(&models.ProductionCompany{})
	//d.DB.AutoMigrate(&models.ProductionCountry{})
	//d.DB.AutoMigrate(&models.MoviesGenre{})
	//d.DB.AutoMigrate(&models.MoviesProductionCompany{})
	//d.DB.AutoMigrate(&models.MoviesKeyword{})

	return nil
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func (d *Database) CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to close connection from database", err)
		return
	}
	dbSQL.Close()
}
