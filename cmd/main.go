package main

import (
	"FractalCompression/internal/compression"
	"FractalCompression/internal/database/postgres"
	"github.com/jackc/pgx"
	"log"
)

func main() {
	//получение конфигураций
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "127.0.0.1",
			Port:     5432,
			Database: "testing",
			User:     "postgres",
			Password: "password",
		},
		MaxConnections: 20,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	})
	log.Print(err)
	db := postgres.NewPG(pool)

	key := compression.Key{Name: "hashCode", Type: "bigint",
		Len: 8, Users: true,
		Script: "CREATE OR REPLACE FUNCTION hashCode(_string text) " +
			"RETURNS BIGINT AS $$\nDECLARE\n  val_ CHAR[];\n  " +
			"h_ BIGINT := 0;\n  ascii_ BIGINT;\n  " +
			"c_ char;\nBEGIN\n " +
			" val_ = regexp_split_to_array(_string, '');\n\n  " +
			"FOR i in 1 .. array_length(val_, 1)\n  " +
			"LOOP\n    c_ := (val_)[i];\n    " +
			"ascii_ := ascii(c_);\n    " +
			"h_ = (31 * h_ + ascii_ ) % (1e9 + 9);\n    " +
			"raise info '%: % = %', i, c_, h_;\n  " +
			"END LOOP;\n" +
			"RETURN h_;\n" +
			"END;\n" +
			"$$ LANGUAGE plpgsql;"}
	tb := compression.NewTable(10, "music", db, key)
	log.Print(tb.Compress())
}
