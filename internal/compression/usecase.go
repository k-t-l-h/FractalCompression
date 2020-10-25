package compression

import (
	ga "github.com/k-t-l-h/GenAlgo"
	"github.com/pkg/errors"
	"log"
	"math"
	"math/rand"
	"strings"
	"sync"
)

func (t *Table) Compress() error {

	log.Print("getting meta")
	if ok := t.getMeta(); ok != nil {
		return ok
	}

	log.Print("getting constraints")
	if ok := t.getConstrains(); ok != nil {
		return ok
	}

	log.Print("getting values")
	if ok := t.getValues(); ok != nil {
		return ok
	}

	log.Print("getting unique values")
	if ok := t.getUniqueValues(); ok != nil {
		return ok
	}

	log.Print("getting data len")
	if ok := t.getDataLen(); ok != nil {
		return ok
	}

	log.Print("getting compressible")
	if ok := t.getCompressible(); ok != nil {
		return ok
	}

	log.Print("getting priorities")
	if ok := t.getPriorities(); ok != nil {
		return ok
	}

	log.Print("getting domens")
	if ok := t.getDomens(); ok != nil {
		return ok
	}

	log.Print("getting compression")
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

	state = false
	for i := 0; i < len(t.Columns); i++ {
		wg.Add(1)
		go func(i int) {
			//ограничения описаны по стилю postgres
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
func (t *Table) getUniqueValues() error {
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
		return errors.New("error while getting unique values")
	}

	return nil

}

//заполнение информации о количестве строк в таблице
//возвращает ошибку, если количество получить не удалось
func (t *Table) getValues() error {
	var wg sync.WaitGroup
	var state bool
	chn := make(chan error, len(t.Columns))

	for i := 0; i < len(t.Columns); i++ {

		wg.Add(1)
		go func(i int) {
			value, err := t.Database.GetValues(t.Name, t.Columns[i].Name)
			t.Columns[i].Values = value
			chn <- err
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < len(t.Columns); i++ {
		state = state || (<-chn != nil)
	}

	if state {
		return errors.New("error while getting values")
	}

	return nil
}

//заполнение информации о длине типов данных
//возвращает ошибку, если длины получить не удалось
func (t *Table) getDataLen() error {
	var wg sync.WaitGroup
	var state bool
	chn := make(chan error, len(t.Columns))

	for i := 0; i < len(t.Columns); i++ {

		wg.Add(1)
		go func(i int) {
			value, err := t.Database.Size(t.Columns[i].Type, 1)
			t.Columns[i].DataLen = value
			chn <- err
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < len(t.Columns); i++ {
		state = state || (<-chn != nil)
	}

	if state {
		return errors.New("error while getting data len")
	}

	return nil
}

//определение сжимаемых и несжимаемых столбцов таблицы
//возвращает ошибку, если нет столбцов для сжатия или их меньше двух
func (t *Table) getCompressible() error {

	for i, col := range t.Columns {

		state := false
		//точно несжимаемые столбцы
		state = state || col.Constrains.Key
		state = state || col.Constrains.Sequence
		state = state || col.Constrains.Exclusion
		state = state || col.Constrains.PrimaryKey
		state = state || col.Constrains.ReferenceKey

		//потенциально несжимаемые
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

	for _, col := range t.Columns {
		col.Priority = 1
	}
	return nil
}

//генетический алгоритм
func (t *Table) getDomens() error {
	var dbError error
	var mapMutex sync.Mutex
	columnValues := make(map[int]float64)

	gao := ga.GenAlgo{
		MaxIteration: 100,
		Generator:    &ga.Generator{Len: len(t.Compressible)},
		Crossover: &ga.NPointCrossover{
			N:               2,
			Probability:     0.7,
			ProbabilityFunc: rand.Float64,
		},
		Mutate: &ga.OneDotMutatation{
			Probability:     1,
			ProbabilityFunc: rand.Float64,
		},
		Schema: &ga.Truncation{},
		Fitness: func(unit ga.BaseUnit) float64 {

			bits := unit.GetCromosomes()
			names := ""
			count := 0
			id := 0
			num := len(bits)
			for i, j := range bits {
				id += int(math.Pow(2, float64(num))) * j
				if j == 1 {
					count++
					names += "\"" + t.Columns[t.Compressible[i]].Name + "\", "
				}
			}
			if count == 0 {
				unit.SetFitness(-1)
				return -1
			} else if count < 2 {
				unit.SetFitness(-1)
				return -1
			}

			names = names[:len(names)-2]

			mapMutex.Lock()
			values := columnValues[id]
			mapMutex.Unlock()

			if values != 0 {
				unit.SetFitness(values)
				return values
			}

			var v uint64
			v, dbError = t.Database.GetUniqueValues(t.Name, names)
			values = (float64(t.Columns[0].Values) - float64(v)) / float64(t.Columns[0].Values)

			var sum uint64
			sum = 0
			//сколько весит одна строчка
			for i := 0; i < len(bits); i++ {
				if bits[i] == 1 {
					v2 := t.Columns[t.Compressible[i]].DataLen * v
					sum += v2
				}
			}

			//ключи в словарной таблице
			//+ключи в главной таблице
			//+n строк
			if (t.key.Len*v + t.key.Len*t.Columns[0].Values + sum*v) > sum*t.Columns[0].Values {
				values = -1
			}

			unit.SetFitness(values)
			mapMutex.Lock()
			columnValues[id] = values
			mapMutex.Unlock()
			return values
		},
		Select: &ga.Panmixia{},
		Exit: func() bool {
			return dbError != nil
		},
	}

	gao.Init(len(t.Compressible) * 10)
	gao.Simulation()

	if dbError != nil {
		return errors.Wrap(dbError, "error while getting domes: ")
	}

	max := gao.Population[0].GetCromosomes()
	for i := 0; i < len(max); i++ {
		if max[i] == 1 {
			t.Domens = append(t.Domens, t.Compressible[i])
		}
	}

	return nil
}

func (t *Table) compressData() error {

	var c []string
	var cd []string
	var u []string

	j := 0
	for i := 0; i < len(t.Domens); {
		if j == t.Domens[i] {
			c = append(c, t.Columns[j].Name)
			cd = append(cd, t.Columns[j].Type)
			i++
		} else {
			u = append(u, t.Columns[j].Name)
		}
		j++
	}

	if t.key.Users {
		err := t.Database.KeyFunction(t.key.Script)
		if err != nil {
			return errors.Wrap(err, "error while data precompressing")
		}
	}

	err := t.Database.PreCompress(c, cd, t.Name, t.key.Name, t.key.Type)
	if err != nil {
		return errors.Wrap(err, "error while data precompressing: ")
	}

	err = t.Database.Compress(c, u, t.Name, t.key.Name)
	if err != nil {
		return errors.Wrap(err, "error while compressing data: ")
	}

	err = t.Database.PostCompress(c, t.Name, t.key.Name, t.key.Type)
	if err != nil {
		return errors.Wrap(err, "error while data post compressing: ")
	}

	return nil
}
