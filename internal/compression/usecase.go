package compression

import (
	"errors"
	"strings"
	"sync"
)

func (t *Table) Compress() {

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
			constrains, err := t.Database.GetConstrain(t.Columns[i].Name)
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

	for i, col := range t.Columns {
		col.Values = value
		if col.UniqueValues > col.Values/t.K {
			t.Incompressible = append(t.Incompressible, i)
		} else {
			t.Compressible = append(t.Compressible, i)
		}
	}
	return nil
}

//получение дерева
