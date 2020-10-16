package internal

type IDatabase interface {
	GetNames(string) ([]string, []string, error)
	GetConstraints(string, string) ([]string, error)
	GetValues(string) (uint64, error)
	GetUniqueValues(string, string) (uint64, error)
	Compress([]string, []string, string) error
}
