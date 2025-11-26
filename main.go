package main

import (
	"log"

	"github.com/archmagejay/excercise_pt/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	startRepl(&cfg)
}