package postgres

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"log"
)

type PG struct {
	pool *pgx.ConnPool
}

func NewPG(pool *pgx.ConnPool) *PG {
	return &PG{pool: pool}
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

//получение количества значений в столбце
func (p *PG) GetValues(tableName string) (uint64, error) {
	var num uint64
	//tableName = pgx.Identifier.Sanitize(tableName)
	GetValue := fmt.Sprintf("SELECT COUNT(*) FROM %s;", tableName)

	row := p.pool.QueryRow(GetValue)
	err := row.Scan(&num)
	if err != nil {
		return 0, errors.Wrap(err, "error while scanning all values: ")
	}
	return num, nil
}

//получение количества уникальных значений в столбце
func (p *PG) GetUniqueValues(tableName string, columnName string) (uint64, error) {
	var num uint64
	//tableName = pgx.Identifier.Sanitize(tableName)
	//columnName = pgx.Identifier.Sanitize(columnName)
	GetUniqueValue := fmt.Sprintf("SELECT COUNT( DISTINCT \"%s\" ) FROM %s;", columnName, tableName)

	row := p.pool.QueryRow(GetUniqueValue)
	err := row.Scan(&num)
	if err != nil {
		return 0, errors.Wrap(err, "error while scanning unique values: ")
	}

	return num, nil
}

//создание таблицы уникальных значений
//добавление в бд таблицы для хранения уникальных значений
func (p *PG) PreCompress(names []string, datatypes []string, tableName string) error {
	data := ""
	unique := ""
	for i, _ := range names {
		//name = pgx.Identifier.Sanitize(name)
		data += "\"" + names[i] + "\" " + datatypes[i] + ", "
		unique += "\"" + names[i] + "\"" + ", "
	}
	data = data[:len(data)-2]
	unique = unique[:len(unique)-2]


	//TODO: перенос ограничений данных
	CreateDatabase := "CREATE TABLE %s_compressed " +
		"( hash text PRIMARY KEY, %s , UNIQUE(%s));"
	CreateDatabase = fmt.Sprintf(CreateDatabase,
		tableName, data, unique)

	_, err := p.pool.Exec(CreateDatabase)
	if err != nil {
		return errors.Wrap(err, "error while creating table: ")
	}
	return nil
}

//вынесение в таблицу уникальных значений
func (p *PG) Compress(compressible []string, other []string, tableName string) error {

	if len(compressible) < 2 {
		return errors.New("error while compressing data: nor enough columns")
	}

	compress := ""
	for _, name := range compressible {
		//name = pgx.Identifier.Sanitize(name)
		compress += "\"" + name + "\"" + ", "
	}

	if len(compress) > 2 {
		compress = compress[:len(compress)-2]
	}

	uncompress := ""
	for _, name := range other {
		//name = pgx.Identifier.Sanitize(name)
		uncompress += "\"" + name + "\"" + ", "
	}
	if len(compress) > 2 {
		uncompress = uncompress[:len(uncompress)-2]
	}

	//TODO: добавить возможность выбора хеш-функции
	compressQuery := fmt.Sprintf("INSERT INTO %s_compressed ("+
		"SELECT md5(ROW(%s)::TEXT), %s FROM %s) ON CONFLICT do nothing",
		tableName, compress, compress, tableName)

	exec, err := p.pool.Exec(compressQuery)
	if err != nil {
		return errors.Wrap(err, "error while compressing data: ")
	} else if exec.RowsAffected() == 0 {
		return errors.New("error while compressing data: nothing was compressed")
	}
	return nil
}

//изменение исходной таблицы
func (p *PG) PostCompress(compressible []string, tableName string) error {

	//TODO: error handling & transaction

	//alter table
	AddKey := fmt.Sprintf("ALTER TABLE \"%s\" ADD COLUMN hash TEXT", tableName)
	log.Print(p.pool.Exec(AddKey))

	//занесение новых значений
	compress := ""
	for _, name := range compressible {
		//name = pgx.Identifier.Sanitize(name)
		compress += "\"" + name + "\"" + ", "
	}

	if len(compress) > 2 {
		compress = compress[:len(compress)-2]
	}

	compressQuery := fmt.Sprintf("UPDATE %s SET hash= md5(row(%s)::TEXT);", tableName, compress)
	log.Print(compressQuery)
	log.Print(p.pool.Exec(compressQuery))

	//сброс старых колонок
	for _, c := range compressible {
		DropKey := fmt.Sprintf("ALTER TABLE %s DROP COLUMN  \"%s\"", tableName, c)
		log.Print(p.pool.Exec(DropKey))
	}

	//установка ограничений
	FKformat := "ALTER TABLE %s ADD CONSTRAINT hash_pkey " +
		"FOREIGN KEY (hash) REFERENCES %s_compressed (hash) NOT VALID"
	ForeignKey := fmt.Sprintf(FKformat, tableName, tableName)
	log.Print(p.pool.Exec(ForeignKey))

	//валидация ограничений
	ForeignKeyValidate := fmt.Sprintf("ALTER TABLE %s VALIDATE CONSTRAINT hash_pkey;", tableName)
	log.Print(p.pool.Exec(ForeignKeyValidate))

	//TODO: vacuum
	p.pool.Exec("VACUUM  wines2;")
	return nil
}

