package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/ziqiongg/hackernews/graph"
	"github.com/ziqiongg/hackernews/graph/generated"
	"github.com/ziqiongg/hackernews/internal/pkg/db"
	//"gorm.io/gorm"
	//db "github.com/ziqiongg/hackernews/internal/pkg/db"
)

const defaultPort = "8080"
// const connStr = "postgresql://carolli:password@hackernews/todos?sslmode=disable"

func main() {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// // conn := db.ConnString("localhost", 5432, "carolli","hackernews")
	// // db.New(conn)
	// // defer db.CloseDB()

	// router := chi.NewRouter()
	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// router.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	// conn := db.ConnString("localhost", 5432, "carolli","hackernews")
	// db.New(conn)
	// defer db.CloseDB()

	newDB,err := db.NewConnection()
	if err != nil{
		panic(err)
	}

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: newDB}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))



	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
