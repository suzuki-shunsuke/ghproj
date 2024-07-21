package add

import (
	"fmt"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

func findConfig(fs afero.Fs) (afero.File, error) {
	f, err := fs.Open("ghproj.yaml")
	if err == nil {
		return f, nil
	}
	return fs.Open(".ghproj.yaml") //nolint:wrapcheck
}

func findAndReadConfig(fs afero.Fs, cfg *Config) error {
	f, err := findConfig(fs)
	if err != nil {
		return fmt.Errorf("find a configuration file: %w", err)
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return fmt.Errorf("decode a configuration file: %w", err)
	}
	return nil
}
