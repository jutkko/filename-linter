package main

import (
	"flag"
	"log"

	"github.com/jutkko/filename-linter/linter"
)

func main() {
	var dir string
	flag.StringVar(&dir, "d", ".", "the directory to run the linter in")
	flag.Parse()

	err := linter.LintFiles(dir)
	if err != nil {
		log.Fatal(err)
	}
}
