package confuso

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func printConfig(config map[string]string) {
	for key, value := range config {
		fmt.Println("'" + key + "': '" + value + "'")
	}
}

func readConfig(fileName string) (map[string]string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var config = make(map[string]string)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
		line = strings.Trim(line, " \n")

		// Skip empty lines and comments
		if line == "" || line[0] == '#' {
			continue
		}

		splitted := strings.Split(line, "=")
		if len(splitted) != 2 {
			return nil, errors.New("malformed line in config: " + line)
		}
		config[splitted[0]] = splitted[1]
	}
	return config, nil
}
