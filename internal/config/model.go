package config

type CompressionConfig struct {
	DC DatabaseConfig `json:"database"`
	TC TableConfig    `json:"table"`
	KC KeyConfig      `json:"key"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     uint   `json:"port"`
	MaxConn  uint   `json:"max_conn"`
}

type TableConfig struct {
	K        uint64 `json:"k"`
	Name     string `json:"name"`
	Strategy string `json:"strategy"`
}

type KeyConfig struct {
	Name   string `json:"name"`
	Users  bool   `json:"users,omitempty"`
	Len    uint64 `json:"len,omitempty"`
	Script string `json:"script,omitempty"`
	Type   string `json:"type"`
}
