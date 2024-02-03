package helper

import (
	"bufio"
	"os"
	"strings"
)

func ReadConfig(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	homeDir := os.Getenv("HOME")

	var files []string
	var excludes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "exclude:") {
			excludes = append(excludes, strings.TrimSpace(strings.TrimPrefix(line, "exclude:")))
		} else {
			filePath := strings.TrimSpace(line)
			if strings.HasPrefix(filePath, "~/.") {
				filePath = strings.Replace(filePath, "~", homeDir, 1)
			}
			if strings.HasPrefix(filePath, "~/") {
				filePath = strings.Replace(filePath, "~", homeDir, 1)
			}

			files = append(files, filePath)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return files, excludes, nil
}
