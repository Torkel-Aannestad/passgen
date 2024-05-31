package internal

import (
	"embed"
	"encoding/csv"
	"io"
)

//go:embed dictionary.csv
var f embed.FS

func LoadWordsFromCSV() ([]string, error) {
	words := []string{}

	file, err := f.Open("dictionary.csv")
	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		words = append(words, record[0])
	}

	return words, nil
}
