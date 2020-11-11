package compression

import (
	combinations "github.com/mxschmitt/golang-combinations"
	"log"
)

func (t *Table) BruteForceStrategy() error {
	//список всех имен
	var names []string
	//получение списка всех имен
	for _, i := range t.Compressible {
		names = append(names, t.Columns[i].Name)
	}
	log.Print(names)
	all := combinations.All(names)

	var best []string
	var bestScore float64

	bestScore = 1
	for _, one := range all {
		str := ""
		for _, name := range one {
			str = str + name + ", "
		}
		str = str[:len(str)-2]

		if len(one) != 1 || len(one) != len(t.Compressible) {
			v, _ := t.Database.GetUniqueValues(t.Name, str)

			values := (float64(t.Columns[0].Values) - float64(v)) / float64(t.Columns[0].Values)
			if values < bestScore {
				best = one
			}
		}
	}


	//TODO: один проход
	for _, b := range best {
		for i, _ := range t.Columns {
			if b == t.Columns[i].Name {
				t.Domens = append(t.Domens, i)
			}
		}
	}
	return nil
}
