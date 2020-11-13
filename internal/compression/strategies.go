package compression

import (
	ga "github.com/k-t-l-h/GenAlgo"
	combinations "github.com/mxschmitt/golang-combinations"
	"github.com/pkg/errors"
	"math"
	"math/rand"
	"sort"
	"sync"
)

type Case struct {
	Max   float64
	Names []string
	Value float64
}

type Cases []Case

func (c Cases) Len() int {
	return len(c)
}

func (c Cases) Less(i int, j int) bool {

	if c[i].Value > c[i].Max && c[j].Value > c[j].Max {
		//оба соответствуют условию
		return len(c[i].Names) < len(c[j].Names)
	} else if c[i].Value < c[i].Max && c[j].Value > c[j].Max {
		//первый больше порога
		return true
	} else if c[i].Value > c[i].Max && c[j].Value < c[j].Max {
		//второй больше порога
		return false
	} else {
		return len(c[i].Names) < len(c[j].Names)
	}
}

func (c Cases) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (t *Table) BruteForceStrategy() error {

	//список всех имен
	var names []string
	//получение списка всех имен
	for _, i := range t.Compressible {
		names = append(names, t.Columns[i].Name)
	}
	all := combinations.All(names)

	var cs []Case

	count := 0
	max := float64(t.Columns[0].Values) - float64(t.Columns[0].Values*(1/t.K))/float64(t.Columns[0].Values)
	for _, one := range all {
		count = 0
		str := ""
		for _, name := range one {
			count++
			str = str + name + ", "
		}
		str = str[:len(str)-2]

		//TODO: придумать формулу
		if count > 1 {
			v, _ := t.Database.GetUniqueValues(t.Name, str)
			values := (float64(t.Columns[0].Values) - float64(v)) / float64(t.Columns[0].Values)
			cs = append(cs, Case{
				Names: one,
				Value: values,
				Max:   max,
			})
		}

	}

	sort.Sort(Cases(cs))
	var best Case

	if len(cs) > 0 {
		best = cs[0]
	}

	var i int
	var b int
	i = 0
	b = 0
	for i = 0; i < len(t.Columns) && b < len(best.Names); i++ {
		if t.Columns[i].Name == best.Names[b] {
			t.Domens = append(t.Domens, i)
			b++
		}
	}
	//TODO: проверки и обработка ошибок
	return nil
}

func (t *Table) GenAlgoStrategy() error {
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
