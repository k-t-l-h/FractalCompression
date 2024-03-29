package internal

type IDatabase interface {
	GetNames(string) ([]string, []string, error)
	GetConstraints(string, string) ([]string, error)
	GetValues(string, string) (uint64, error)
	GetUniqueValues(string, string) (uint64, error)
	PreCompress([]string, []string, string, string, string) error
	Compress([]string, []string, string, string) error
	PostCompress([]string, string, string, string) error
	KeyFunction(string) error
	Size(string, uint64) (uint64, error)
}
