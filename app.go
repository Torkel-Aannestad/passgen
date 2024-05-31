package main

import (
	"flag"
	"fmt"
)

func app() {
	fmt.Println("From app")
	type config struct {
		words          *int
		includeNumbers *bool
		capitalize     *bool
		separator      *string
	}
	type passphrase struct {
		config *config
	}

	cfg := config{
		words:          flag.Int("w", 4, "number of words in the passphrase"),
		includeNumbers: flag.Bool("n", true, "include numbers in the passphrase"),
		capitalize:     flag.Bool("c", true, "capitalize the first letter of each word"),
		separator:      flag.String("s", " ", "character to separate words in the passphrase"),
	}
	flag.Parse()
	fmt.Printf("words: %v", cfg.words)
}
