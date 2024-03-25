package main

import (
	"flag"
	"os"
)

func main()  {
	var phrase string

	flag.StringVar(&phrase, "phrase", "Enter a phrase", "A phrase to determine the letter/word frequency")
	flag.Parse()

	if phrase == " "{
		os.Exit(1)
	}

	
}