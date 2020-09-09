package ethutils

import (
	"bufio"
	"os"
)

// MustReadLines returns all lines in file.
func MustReadLines(f string) []string {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := make([]string, 0)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return list
}
