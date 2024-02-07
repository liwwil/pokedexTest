package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TanyaEIEI/pokedex/common"
	"github.com/TanyaEIEI/pokedex/graph"
	"github.com/go-chi/chi"
)

func main() {
	var err error
	port := "9000"
	if err != nil {
		log.Fatal(err)
	}
	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	customCtx := graph.CustomContext{
		Database: db,
	}

	r := chi.NewRouter()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", graph.CreateContext(&customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
