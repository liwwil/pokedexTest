package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TanyaEIEI/pokedex/database"
	"github.com/TanyaEIEI/pokedex/graph"
	"github.com/go-chi/chi"
)

const defaultPort = "19000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, _ := database.InitDb()
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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
