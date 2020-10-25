package compression

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
