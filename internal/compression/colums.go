package compression

type Column struct {
	Name       string
	Type       string
	Constrains struct {
		//несжимаемость
		PrimaryKey bool
		Key        bool
		Exclusion  bool
		Sequence   bool
		//потенциальная несжимаемость
		ReferenceKey bool
		Users        bool
	}

	Values       uint64
	UniqueValues uint64
	Priority uint64
}
