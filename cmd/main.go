package main

import (
	"FractalCompression/internal/compression"
	"FractalCompression/internal/config"
	"FractalCompression/internal/database/postgres"
	"log"
)

func main() {

	cnf, err := config.GetData("data.json")
	if err != nil {
		log.Print(err)
		return
	}

	db, err := postgres.NewPG(cnf.DC)
	if err != nil {
		log.Print(err)
		return
	}

	key := compression.NewKey(cnf.KC)
	tb := compression.NewTable(&cnf.TC, db, key)
	log.Print(tb.Compress())
}
