package confuso

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLInput struct {
	filePath string
}

func NewYAMLInput(filePath string) Input {
	return &YAMLInput{filePath: filePath}
}

func (y *YAMLInput) read() (map[string]any, error) {
	content, err := os.ReadFile(y.filePath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var out map[string]any

	if err := yaml.Unmarshal(content, &out); err != nil {
		return nil, fmt.Errorf("parsing YAML: %w", err)
	}

	return out, nil
}
