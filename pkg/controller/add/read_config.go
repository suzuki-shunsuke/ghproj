package add

import (
	"fmt"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

func findConfig(fs afero.Fs, p string) (afero.File, error) {
	if p != "" {
		f, err := fs.Open(p)
		if err != nil {
			return nil, fmt.Errorf("open a configuration file: %w", err)
		}
		return f, nil
	}
	f, err := fs.Open("ghproj.yaml")
	if err == nil {
		return f, nil
	}
	return fs.Open(".ghproj.yaml") //nolint:wrapcheck
}

func findAndReadConfig(fs afero.Fs, cfg *Config, param *Param) error {
	if param.ConfigText != "" {
		if err := yaml.Unmarshal([]byte(param.ConfigText), cfg); err != nil {
			return fmt.Errorf("unmarshal a configuration text: %w", err)
		}
		return nil
	}
	f, err := findConfig(fs, param.ConfigFilePath)
	if err != nil {
		return fmt.Errorf("find a configuration file: %w", err)
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return fmt.Errorf("decode a configuration file: %w", err)
	}
	return nil
}
