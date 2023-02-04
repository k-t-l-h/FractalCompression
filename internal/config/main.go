package config

import (
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
	"os"
)

func GetData(filename string) (CompressionConfig, error) {
	var config CompressionConfig
	reader, err := os.Open(filename)
	if err != nil {
		return CompressionConfig{}, errors.Wrap(err, "error while file opening: ")
	}
	err = easyjson.UnmarshalFromReader(reader, &config)
	if err != nil {
		return CompressionConfig{}, errors.Wrap(err, "error while getting config dara: ")
	}
	return config, nil
}
