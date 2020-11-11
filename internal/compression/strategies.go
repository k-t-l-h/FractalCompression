package compression

import (
	combinations "github.com/mxschmitt/golang-combinations"
	"math"
)

func (t *Table) BruteForceStrategy() error {
	//список всех имен
	names := []string{"A", "B", "C"}
	//получение списка всех имен
	for _, i := range t.Compressible {
		names = append(names, t.Columns[i].Name)
	}

	all := combinations.All(names)
	var best []string
	var bestscore uint64
	str := ""
	bestscore = math.MaxUint64
	for _, one := range all {
		for _, name := range one {
			str += name + ", "
		}
		str = str[:len(str)-2]
		v, _ := t.Database.GetUniqueValues(t.Name, str)
		if v < bestscore {
			best = one
		}
	}

	for _, b := range best {
		for i, _ := range t.Columns {
			if b == t.Columns[i].Name {
				t.Domens = append(t.Domens, i)
			}
		}
	}

	return nil
}
