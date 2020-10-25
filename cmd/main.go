package main

import (
	"FractalCompression/internal/compression"
	"FractalCompression/internal/config"
	"FractalCompression/internal/database/postgres"
	"github.com/jackc/pgx"
	"log"
)

func main() {

	cnf, err := config.GetData("data.json")
	if err != nil {
		log.Print(err)
		return
	}

	//получение конфигураций
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     cnf.DC.Host,
			Port:     uint16(cnf.DC.Port),
			Database: cnf.DC.Database,
			User:     cnf.DC.User,
			Password: cnf.DC.Password,
		},
		MaxConnections: int(cnf.DC.MaxConn),
		AfterConnect:   nil,
		AcquireTimeout: 0,
	})
	if err != nil {
		log.Print(err)
		return
	}
	db := postgres.NewPG(pool)

	key := compression.Key{Name: cnf.KC.Name, Type: cnf.KC.Type,
		Len: cnf.KC.Len, Users: cnf.KC.Users,
		Script: cnf.KC.Script}
	tb := compression.NewTable(cnf.TC.K, cnf.TC.Name, db, key)
	log.Print(tb.Compress())
}
