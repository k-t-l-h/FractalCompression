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
	cpuprofile string
	memprofile string
)

func init() {
	flag.StringVar(&fn, "filename", "", "set the filename of config")
	//flag.String("cpuprofile", "cpu.prof", "write cpu profile to `file`")
	//flag.String("memprofile", "mem.prof", "write memory profile to `file`")
}

func main() {

	flag.Parse()
	/*
	cpuprofile = "cpu.prof"
	memprofile = "mem.prof"
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	*/


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

	/*
	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	*/
}
