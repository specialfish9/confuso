package input

import (
	"fmt"
	"os"

	"github.com/specialfish9/confuso/v2"
	yaml3 "gopkg.in/yaml.v3"
)

type yaml struct {
	filePath string
}

func YAML(filePath string) confuso.Input {
	return &yaml{filePath: filePath}
}

func (y *yaml) Read() (map[string]any, error) {
	content, err := os.ReadFile(y.filePath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var out map[string]any

	if err := yaml3.Unmarshal(content, &out); err != nil {
		return nil, fmt.Errorf("parsing YAML: %w", err)
	}

	return out, nil
}
