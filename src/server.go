package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dreamingkills/steep/graph"
	"github.com/dreamingkills/steep/graph/generated"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	db, err := pgxpool.Connect(context.Background(), os.Getenv("DSN"))
	if err != nil {
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
