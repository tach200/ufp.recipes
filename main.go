package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

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

	r := mux.NewRouter()

	r.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
	r.HandleFunc("/recipe/{id}", handlers.GetRecipe).Methods("GET")

	http.Handle("/", r)

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

	fmt.Print("connected to db\n")

	handlers := &handlers.Handlers{
		DB: db,
	}

	return *handlers
}
