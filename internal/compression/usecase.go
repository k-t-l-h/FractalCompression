package compression

import (
	"errors"
	"log"
	"strings"
	"sync"
)

func (t *Table) Compress() error {

	if ok := t.getMeta(); ok != nil {
		return ok
	}

	if ok := t.getConstrains(); ok != nil {
		return ok
	}

	if ok := t.getValue(); ok != nil {
		return ok
	}

	if ok := t.getValueFactor(); ok != nil {
		return ok
	}

	if ok := t.getCompressible(); ok != nil {
		return ok
	}

	if ok := t.getPriorities(); ok != nil {
		return ok
	}

	if ok := t.getDomens(); ok != nil {
		return ok
	}

	if ok := t.compressData(); ok != nil {
		return ok
	}

	return nil
}

//получение названий и типов столбцов таблицы
//возвращает ошибку, если названия получить не удалось
func (t *Table) getMeta() error {
	//заполнение информации о столбцах
	names, types, err := t.Database.GetNames(t.Name)

	if err != nil {
		return errors.New("database error while parsing names")
	}

	//создание массива столбцов
	t.Columns = make([]*Column, len(types), len(types))

	//заполнение информации о столбцах базы данных
	for i := 0; i < len(names); i++ {
		c := Column{Name: names[i], Type: types[i]}
		t.Columns[i] = &c
	}

	return nil
}

//получение ограничений, наложенных на столбцы
//возвращает ошибку, если ограничения получить не удалось
func (t *Table) getConstrains() error {

	var wg sync.WaitGroup
	var state bool
	chn := make(chan error, len(t.Columns))

	for i := 0; i < len(t.Columns); i++ {
		wg.Add(1)
		go func(i int) {
			constrains, err := t.Database.GetConstraints(t.Name, t.Columns[i].Name)
			for _, name := range constrains {
				if strings.HasSuffix(name, "_pkey") {
					t.Columns[i].Constrains.PrimaryKey = true
				} else if strings.HasSuffix(name, "_excl") {
					t.Columns[i].Constrains.Exclusion = true
				} else if strings.HasSuffix(name, "_seq") {
					t.Columns[i].Constrains.Sequence = true
				} else if strings.HasSuffix(name, "_key") {
					t.Columns[i].Constrains.Key = true
				} else if strings.HasSuffix(name, "_fkey") {
					t.Columns[i].Constrains.Key = true
				} else {
					t.Columns[i].Constrains.Users = true
				}
			}
			chn <- err
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < len(t.Columns); i++ {
		state = state || (<-chn != nil)
	}

	if state {
		return errors.New("error while getting constrains")
	}

	return nil
}

//заполнение информации о количестве уникальных значений в столбцах
//возвращает ошибку, если количество получить не удалось
func (t *Table) getValueFactor() error {
	var wg sync.WaitGroup
	var state bool
	chn := make(chan error, len(t.Columns))

	for i := 0; i < len(t.Columns); i++ {

		wg.Add(1)
		go func(i int) {
			value, err := t.Database.GetUniqueValues(t.Name, t.Columns[i].Name)
			t.Columns[i].UniqueValues = value
			chn <- err
			defer wg.Done()
		}(i)

	}

	wg.Wait()

	for i := 0; i < len(t.Columns); i++ {
		state = state || (<-chn != nil)
	}

	if state {
		return errors.New("error while getting constrains")
	}

	return nil

}

//заполнение информации о количестве строк в таблице
//заполнение информации о индексах сжимаемых и несжимаемых столбцов
//возвращает ошибку, если количество получить не удалось
func (t *Table) getValue() error {
	value, err := t.Database.GetValues(t.Name)
	if err != nil {
		return errors.New("error while counting values")
	}

	for _, col := range t.Columns {
		col.Values = value
	}
	return nil
}

//определение сжимаемых и несжимаемых столбцов таблицы
//возвращает ошибку, если нет столбцов для сжатия или их меньше двух
func (t *Table) getCompressible() error {

	for i, col := range t.Columns {
		//
		state := false
		//точно несжимаемые столбцы
		state = state || col.Constrains.Key
		state = state || col.Constrains.Sequence
		state = state || col.Constrains.Exclusion
		state = state || col.Constrains.PrimaryKey
		state = state || col.Constrains.ReferenceKey

		//потенциально несжимаемые
		//TODO: обработка пользовательских ограничений

		if !state && (col.UniqueValues <= col.Values/t.K) {
			t.Compressible = append(t.Compressible, i)
		} else {
			t.Incompressible = append(t.Incompressible, i)
		}
	}

	if len(t.Compressible) < 2 {
		return errors.New("compression is not possible")
	}
	return nil
}

//определение приоритетов столбцов
func (t *Table) getPriorities() error {

	//TODO: получение дерева приоритетов из базы данных
	for _, col := range t.Columns {
		//TODO: подстановка приоритета на основе дерева приоритетов
		col.Priority = 1
	}
	return nil
}

//генетический алгоритм
func (t *Table) getDomens() error {
	//TODO: генетический алгоритм
	//TODO: обработка ошибок связи с бд внутри генетического алгоритма
	t.Domens = []int{1, 2}
	return nil
}

func (t *Table) compressData() error {
	//TODO: получение имен столбцов
	c := []string{"valueA", "valueB"}
	u := []string{"data"}

	log.Print(t.Database.PreCompress(c, []string{"integer", "integer"}, t.Name))

	err := t.Database.Compress(c, u, t.Name)
	if err != nil {
		return errors.New("error while compressing data")
	}
	return nil
}
