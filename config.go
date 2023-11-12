package confuso

import (
	"bufio"
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
		splitted := strings.Split(strings.Trim(line, " \n"), "=")
		config[splitted[0]] = splitted[1]
	}
	return config, nil
}
