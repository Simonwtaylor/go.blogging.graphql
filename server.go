package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Simonwtaylor/blogging-gql/entities"
	"github.com/Simonwtaylor/blogging-gql/graph"
	"github.com/Simonwtaylor/blogging-gql/graph/generated"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const defaultPort = "8080"

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=54320 user=postgres dbname=blogging password=password sslmode=disable")

	if err != nil {
		log.Panicf("unable to open postgres db %v", err)
	}

	db.AutoMigrate(&entities.User{})
	db.Create(&entities.User{Email: "nickdunn@gmail.com", Password: "Woop", Username: "Nick"})

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
