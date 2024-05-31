package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"
)

func loadWordsFromCSV(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var words []string
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

func generateWord() string {
	csvFilePath := "./dictionary.csv"
	dictionary, err := loadWordsFromCSV(csvFilePath)
	if err != nil {
		fmt.Println("Error loading words from CSV:", err)
	}

	return dictionary[rand.Intn(len(dictionary))]

}

func generateNumber() string {
	return fmt.Sprintf("%d", rand.Intn(10))
}

func generatePassphrase(numWords int, includeNumbers, capitalize bool, separator string) string {

	var passphrase []string
	numberPosition := rand.Intn(numWords)

	for i := 0; i < numWords; i++ {
		word := generateWord()
		if capitalize {
			firstRune, width := utf8.DecodeRuneInString(word)
			firstCharUpper := strings.ToUpper(string(firstRune))
			word = firstCharUpper + word[width:]
		}
		if includeNumbers == true && i == numberPosition {
			word += generateNumber()
		}
		passphrase = append(passphrase, word)
	}

	return strings.Join(passphrase, separator)
}
func writeToClipboardWindows(text string) error {
	cmd := exec.Command("clip")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func main() {

	words := flag.Int("w", 4, "number of words in the passphrase")
	includeNumbers := flag.Bool("n", true, "include numbers in the passphrase")
	capitalize := flag.Bool("c", true, "capitalize the first letter of each word")
	separator := flag.String("s", " ", "character to separate words in the passphrase")

	flag.Parse()

	passphrase := generatePassphrase(*words, *includeNumbers, *capitalize, *separator)
	writeToClipboardWindows(passphrase)

	fmt.Println(passphrase)
	app()
}
