package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/tach200/ufp.recipes/internal/db"
	"github.com/tach200/ufp.recipes/internal/handlers"
)

var (
	port     string
	dbDriver string
	dbOpts   string
)

func init() {
	flag.StringVar(&port, "port", ":8080", "port number that the server runs on")
	flag.StringVar(&dbDriver, "db-driver", "postgres", "the type of db driver which is being used")
	flag.StringVar(
		&dbOpts,
		"db-opts",
		"host=localhost port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable",
		"database options",
	)
	flag.Parse()
}

func main() {
	handlers := newHandlers(dbDriver, dbOpts)

	router := mux.NewRouter()

	// Create a CORS handler with the desired options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Replace with your React app's domain
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)

	http.Handle("/", handler)

	router.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
	router.HandleFunc("/recipe/{id}", handlers.GetRecipe).Methods("GET")
	router.HandleFunc("/recipe/{id}/products", handlers.GetRecipeProducts).Methods("GET")
	router.HandleFunc("/test", handlers.Test).Methods("GET")

	fmt.Printf("INFO: starting recipe service on port %s\n", port)
	http.ListenAndServe(port, nil)
}

func newHandlers(dbDriver, dbOpts string) handlers.Handlers {
	db, err := db.NewDB(dbDriver, dbOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Client.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("INFO: connected to db\n")

	handlers := &handlers.Handlers{
		DB: db,
	}

	return *handlers
}
