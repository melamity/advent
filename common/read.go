package common

import (
	"bufio"
	"os"
)

func ReadFile(filename string) (<-chan string, error) {
	ch := make(chan string)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}

		defer file.Close()
		close(ch)
	}()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ch, nil
}
