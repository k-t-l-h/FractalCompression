package postgres

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

type PG struct {
	pool *pgx.ConnPool
}

//получение названий столбцов и их типов данных из информационной таблицы базы данных
func (p *PG) GetNames(tableName string) ([]string, []string, error) {
	GetColumnsName := "SELECT column_name, data_type FROM information_schema.columns WHERE table_name = $1"

	rows, err := p.pool.Query(GetColumnsName, tableName)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error while getting columns names: ")
	}

	var names []string
	var types []string

	for rows.Next() {
		var name string
		var tp string
		if ok := rows.Scan(&name, &tp); ok != nil {
			return nil, nil, errors.Wrap(ok, "error while scanning columns names: ")
		}
		names = append(names, name)
		types = append(types, tp)
	}

	return names, types, nil
}

//получение ограничений
func (p *PG) GetConstraints(tableName, columnName string) ([]string, error) {
	//TODO: batch
	GetConstraints := "SELECT constraint_name FROM information_schema.constraint_column_usage " +
		"WHERE table_name = $1 AND column_name = $2"

	rows, err := p.pool.Query(GetConstraints, tableName, columnName)
	if err != nil {
		return nil, errors.Wrap(err, "error while getting constraints names: ")
	}

	var constraints []string

	for rows.Next() {
		var cons string
		if ok := rows.Scan(&cons); ok != nil {
			return nil, errors.Wrap(ok, "error while scanning constraints names: ")
		}

		constraints = append(constraints, cons)
	}

	return constraints, nil
}

//получение уникальных значений в столбце
func (p *PG) GetValues(tableName string) (uint64, error) {
	var num uint64
	tableName = pgx.Identifier.Sanitize(tableName)
	GetValue := fmt.Sprintf("SELECT COUNT(*) FROM %s;", tableName)

	row := p.pool.QueryRow(GetValue)
	err := row.Scan(&num)
	if err != nil {
		return 0, errors.Wrap(err, "error while scanning all values: ")
	}
	return num, nil
}

func (p *PG) GetUniqueValues(tableName string, columnName string) (uint64, error) {
	var num uint64
	tableName = pgx.Identifier.Sanitize(tableName)
	columnName = pgx.Identifier.Sanitize(columnName)

	GetUniqueValue := fmt.Sprintf("SELECT COUNT(DISTINCT %s) FROM %s;", columnName, tableName)

	row := p.pool.QueryRow(GetUniqueValue)
	err := row.Scan(&num)
	if err != nil {
		return 0, errors.Wrap(err, "error while scanning unique values: ")
	}

	return num, nil
}
