package compression

import "FractalCompression/internal"

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
}

func NewTable(k uint64, name string, database internal.IDatabase, key Key) *Table {
	return &Table{K: k, Name: name, Database: database, key: key}
}
