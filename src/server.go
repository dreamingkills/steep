package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dreamingkills/steep/graph"
	"github.com/dreamingkills/steep/graph/generated"
	"github.com/dreamingkills/steep/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load("../.env")
	if(err != nil) {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if(err != nil) {
		panic(err)
	}

	db.Exec("CREATE DATABASE steep;")
	if(db.Error != nil) {
		fmt.Println("unable to create database, attempting to connect...")
	}

	db, err = gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if(err != nil) {
		panic(err)
	}
	
	err = db.AutoMigrate(models.Models...)
	if(err != nil) {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
