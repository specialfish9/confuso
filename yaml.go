package confuso

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func readYAML(filePath string) (map[string]any, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var out map[string]any

	if err := yaml.Unmarshal(content, &out); err != nil {
		return nil, fmt.Errorf("parsing YAML: %w", err)
	}

	return out, nil
}
