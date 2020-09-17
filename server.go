package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/francismarcus/entgql/ent"
	"github.com/francismarcus/entgql/graph"
	"github.com/francismarcus/entgql/graph/generated"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

var client *ent.Client

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := ent.Open("postgres", "postgres://marcusmagnusson@localhost/entgql?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		graph.New(client),
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
