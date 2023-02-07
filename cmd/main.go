package main

import (
	"FractalCompression/internal/compression"
	"FractalCompression/internal/config"
	"FractalCompression/internal/database/postgres"
	"flag"
	"log"
)

var (
	fn string
)

func init() {
	flag.StringVar(&fn, "filename", "", "set the filename of config")
}

func main() {

	flag.Parse()

	cnf, err := config.GetData(fn)
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
