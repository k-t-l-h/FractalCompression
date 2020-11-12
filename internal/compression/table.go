package compression

import (
	"FractalCompression/internal"
	"FractalCompression/internal/config"
)

type Table struct {
	//коэффициент минимального сжатия таблицы
	K uint64
	//название таблицы (необходимо для скриптов)
	Name string
	//таблица для сжатия
	Database internal.IDatabase
	//информация о столбцах таблицы
	Columns []*Column
	//информация о сжимаемых столбцах таблицы
	Compressible []int
	//информация о несжимаемых столбцах таблицы
	Incompressible []int
	//информация о выбранных доменах для сжатия
	Domens []int
	//информация о хеше для кодирования
	key Key
	//информация о стратегии выбора доменов
	Strategy string
}

func NewTable(cnf *config.TableConfig, database internal.IDatabase, key *Key) *Table {
	return &Table{K: cnf.K, Name: cnf.Name, Database: database, key: *key, Strategy: cnf.Strategy}
}
