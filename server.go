package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TanyaEIEI/pokedex/database"
	"github.com/TanyaEIEI/pokedex/graph"
	"github.com/go-chi/chi"
)

const defaultPort = "30001"

func main() {

	db, err := database.InitDb()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	if db == nil {
		log.Fatalf("Error initializing database: %v", db)

	}
	r := chi.NewRouter()
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					Pokedex: &database.Database{
						Db: db,
					},
				},
			},
		),
	)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://192.168.49.2:%s/ for GraphQL playground eiei run passed Yeahhhh", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
