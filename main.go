package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/archmagejay/excercise_pt/internal/config"
	"github.com/archmagejay/excercise_pt/internal/database"
	_ "github.com/lib/pq"
)
type state struct {
	db *database.Queries
	cfg *config.Config
	in  *bufio.Scanner
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
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
		db: dbQueries,
		cfg: &cfg,
		in:  bufio.NewScanner(os.Stdin),
	}

	startRepl(s)
}