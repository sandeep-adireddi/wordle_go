/*
package worddata read words from various input file types into slice 
*/
package worddata

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Read_txt(filename string) ([]string, error) {

	if !strings.HasSuffix(filename, ".txt") {
		return nil, errors.New("not a txt file")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("file not found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())

	}

	return words, nil
}
