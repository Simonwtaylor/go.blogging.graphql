package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Simonwtaylor/blogging-gql/entities"
	"github.com/Simonwtaylor/blogging-gql/graph"
	"github.com/Simonwtaylor/blogging-gql/graph/generated"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Panicf("Error loading .env file, %v", err)
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Panicf("unable to open postgres db %v", err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Post{})
	db.Create(&entities.Post{
		Content:    "Hello world",
		DatePosted: time.Now(),
		User:       entities.User{Email: "nickdunn@gmail.com", Password: "Woop", Username: "Nick"},
	})

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
