package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/shoanchikato/demo/hex/database/sqlite"
	"github.com/shoanchikato/demo/hex/toy"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Getting the database name to use from terminal/commandline
	dbType := flag.String("database", "sqlite", "database type [sqlite, psql]")
	flag.Parse()

	var toyRepo toy.Repository

	switch *dbType {
	case "sqlite":
		sconn := sqliteConnection("file::memory:?cache=shared")
		defer sconn.Close()
		toyRepo = sqlite.NewSqliteToyRepository(sconn)
	case "psql":
		fmt.Println("Not yet Implemented")
		os.Exit(1)
	default:
		fmt.Println("Unknown database")
		os.Exit(1)
	}

	// Injecting services to the handler
	toyService := toy.NewToyService(toyRepo)
	toyHandler := toy.NewToyHandler(toyService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/toys", toyHandler.Get).Methods("GET")
	router.HandleFunc("/toys/{id}", toyHandler.GetByID).Methods("GET")
	router.HandleFunc("/toys", toyHandler.Create).Methods("POST")

	// Enabling CORS to all * origins
	http.Handle("/", accessControl(router))

	// Handling multiple connections and graceful shutdown
	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :3000")
		errs <- http.ListenAndServe(":3000", nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-errs)
}

// Connecting to the sqlite database (function can also be used for postgres by just changing the name and println message)
func sqliteConnection(database string) *sql.DB {
	fmt.Println("Connecting to Sqlite DB")
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatalf("%s", err)
		panic(err)
	}
	
	// Exec Schema for creating toy table
	_, err = db.Exec(sqlite.Schema)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// Function enabling CORS
func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
