package main

import (
	"fmt"
	"log"
	"os"

	srv "github.com/Murolando/hakaton_geo"
	"github.com/Murolando/hakaton_geo/pkg/handler"
	"github.com/Murolando/hakaton_geo/pkg/repository"
	"github.com/Murolando/hakaton_geo/pkg/repository/postgres"
	"github.com/Murolando/hakaton_geo/pkg/service"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env variables:", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPas := os.Getenv("DB_PASSWORD")
	serverPort := os.Getenv("SERVER_PORT")
	dbConfig := postgres.NewConfig(dbHost, dbPort, dbUserName, dbPas, dbName)
	fmt.Println(dbConfig)
	db, err := postgres.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	s := new(srv.Server)
	if err := s.Run(serverPort, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

	fmt.Println('o')
}