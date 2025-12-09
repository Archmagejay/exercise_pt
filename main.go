package main

import (
	"bufio"
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"github.com/archmagejay/excercise_pt/internal/config"
	"github.com/archmagejay/excercise_pt/internal/database"

	_ "github.com/lib/pq"
)

var ErrArgs = errors.New("invalid args")
var ErrNotImplemented = errors.New("not yet implemented")

type state struct {
	db  *database.Queries
	cfg *config.Config
	in  *bufio.Scanner
	l   *log.Logger
}

// Program startup tasks
func main() {
	// Try reading the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Try validating the config file
	if err := cfg.Validate(); err != nil {
		switch err {
		case config.ErrMissingUser:

		case config.ErrDBURL:
			log.Fatalln("Error with database connection")
		case config.ErrTime:
		default:
			log.Fatalf("unknown error validating config: %v", err)
		}
	}
	cfg.LastOpened = time.Now()

	// Attempt to connect to the database
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	// Initialize the logger
	l, err := InitLogger()
	if err != nil {
		log.Fatalf("error setting up logger")
	}

	s := &state{
		db:  dbQueries,
		cfg: cfg,
		in:  bufio.NewScanner(os.Stdin),
		l: l,
	}

	// If everything in initalized start the REPL interface
	startRepl(s)
}


// Shut down the program gracefully
func shutdown(s *state) {
	s.Log(LogInfo, "Closing tracker... Goodbye!")

	// Try to save the config file to disk
	if err := s.cfg.SaveConfig(); err != nil {
		s.Log(LogError, err)
	}

	os.Exit(0)
}
