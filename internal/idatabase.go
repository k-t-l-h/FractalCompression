package internal

type IDatabase interface {
	GetNames(string) ([]string, []string, error)
	GetConstrain(string) ([]string, error)
	GetValues(string) (uint64, error)
	GetUniqueValues(string, string) (uint64, error)
}
