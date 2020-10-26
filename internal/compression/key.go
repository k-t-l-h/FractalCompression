package compression

import "FractalCompression/internal/config"

type Key struct {
	//название хеша для использования в запросах
	Name string
	//тип ключа
	Type string
	//длина ключа
	Len uint64
	//встроенный в бд хеш или нет
	Users bool
	//для не-встроенных: текст скрипта создания
	Script string
}

func NewKey(cnf config.KeyConfig) *Key {
	return &Key{Name: cnf.Name,
		Type:   cnf.Type,
		Len:    cnf.Len,
		Script: cnf.Script}
}
