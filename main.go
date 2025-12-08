package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/archmagejay/excercise_pt/internal/config"
	"github.com/archmagejay/excercise_pt/internal/database"
	"github.com/archmagejay/excercise_pt/internal/logger"
	_ "github.com/lib/pq"
)

var ErrArgs = errors.New("invalid args")

type state struct {
	db  *database.Queries
	cfg *config.Config
	in  *bufio.Scanner
	log *log.Logger
}

func main() {
	l, err := logger.InitLogger()
	if err != nil {
		logger.Log(l, logger.LogFatal, err)
	}

	cfg, err := config.Read()
	if err != nil {
		logger.Log(l, logger.LogFatal, fmt.Sprintf("error reading config: %v", err))
	}

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

	//log.Print(cfg.LastOpened.Format("02/01/2006 15:04:05"))
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: cfg,
		in:  bufio.NewScanner(os.Stdin),
		log: l,
	}

	startRepl(s)
}
