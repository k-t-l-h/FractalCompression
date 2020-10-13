package postgres

import (
	"github.com/jackc/pgx"
)

type PG struct {
	pool *pgx.ConnPool
}

func (p *PG) GetNames(string) ([]string, []string, error) {
	return nil, nil, nil
}

func (p *PG) GetConstrain(string) ([]string, error) {
	return nil, nil
}

func (p *PG) GetValues(string) (uint64, error) {
	return 0, nil
}

func (p *PG) GetUniqueValues(string, string) (uint64, error) {
	return 0, nil
}
