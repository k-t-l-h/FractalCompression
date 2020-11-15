package main

import (
	"FractalCompression/internal/config"
	mg "FractalCompression/internal/database/mongodb"
	"flag"
	"log"
)

var fn string

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

	db, err := mg.NewMG(cnf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(db)
	log.Print(db.GetNames("things"))

	//key := compression.NewKey(cnf.KC)
	//tb := compression.NewTable(&cnf.TC, db, key)
	//log.Print(tb.Compress())
}
