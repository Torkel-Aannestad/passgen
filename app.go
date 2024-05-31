package main

import (
	"flag"
	"fmt"
	"math/rand"
	"passgen/internal"
	"strings"
	"unicode/utf8"
)

type config struct {
	words          *int
	includeNumbers *bool
	capitalize     *bool
	separator      *string
}

type passphrase struct {
	passphrase string
	dictionary []string
	config     config
}

func (p passphrase) generatePassphrase(cfg *config) string {
	var passphrase []string
	numberPosition := rand.Intn(*cfg.words)

	dictionary, err := internal.LoadWordsFromCSV()
	p.dictionary = dictionary
	if err != nil {
		fmt.Println("Error loading words from CSV:", err)
	}

	for i := 0; i < *cfg.words; i++ {
		word := p.generateWord()
		if *cfg.capitalize {
			firstRune, width := utf8.DecodeRuneInString(word)
			firstCharUpper := strings.ToUpper(string(firstRune))
			word = firstCharUpper + word[width:]
		}
		if *cfg.includeNumbers && i == numberPosition {
			word += p.generateNumber()
		}
		passphrase = append(passphrase, word)
	}
	return strings.Join(passphrase, *cfg.separator)
}
func (p passphrase) generateNumber() string {
	return fmt.Sprintf("%d", rand.Intn(10))
}
func (p passphrase) generateWord() string {
	return p.dictionary[rand.Intn(len(p.dictionary))]
}

func app() {

	words := flag.Int("w", 4, "number of words in the passphrase")
	includeNumbers := flag.Bool("n", true, "include numbers in the passphrase")
	capitalize := flag.Bool("c", true, "capitalize the first letter of each word")
	separator := flag.String("s", " ", "character to separate words in the passphrase")

	flag.Parse()

	passphrase := passphrase{
		passphrase: "",
		dictionary: []string{},
		config: config{
			words:          words,
			includeNumbers: includeNumbers,
			capitalize:     capitalize,
			separator:      separator,
		},
	}

	passphrase.passphrase = passphrase.generatePassphrase(&passphrase.config)

	internal.WriteToClipboardWindows(passphrase.passphrase)
	fmt.Println(passphrase.passphrase)

}
