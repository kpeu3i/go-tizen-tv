package samsung

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type TVManagerConfigStorageYAML struct {
	filename string
}

func NewTVManagerConfigStorageYAML(filename string) *TVManagerConfigStorageYAML {
	return &TVManagerConfigStorageYAML{filename: filename}
}

func (s *TVManagerConfigStorageYAML) Load() (TVManagerConfig, error) {
	file, err := os.OpenFile(s.filename, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return TVManagerConfig{}, err
	}

	defer func() {
		_ = file.Close()
	}()

	config := TVManagerConfig{}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil && err != io.EOF {
		return TVManagerConfig{}, err
	}

	return config, nil
}

func (s *TVManagerConfigStorageYAML) Store(config TVManagerConfig) error {
	file, err := os.OpenFile(s.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewEncoder(file)
	err = decoder.Encode(config)
	if err != nil {
		return err
	}

	return nil

}
