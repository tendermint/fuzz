package main

import (
	"flag"
	"os"
	"path/filepath"
)

func main() {
	root := flag.StringVar("root", ".", "the root directory to place the corpus directory in")
	flag.Parse()
	initCorpus(*root)
}

func initCorpus(root string) error {
	corpusDir := filepath.Join(root, "corpus")
	if err := os.MkdirAll(corpusDir); err != nil {
		return err
	}
return nil
}
