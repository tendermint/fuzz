package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	root := flag.String("root", ".", "the root directory to place the corpus directory in")
	flag.Parse()
	initCorpus(*root)
}

func initCorpus(root string) error {
	corpusDir := filepath.Join(root, "corpus")
	if err := os.MkdirAll(corpusDir, 0755); err != nil {
		return err
	}

	corpora := []string{
		"",
		"deadbeef",
		"miscellaneous",
		"A",
		"tendermint tnimrednet",
	}

	for i, corpus := range corpora {
		fullPath := filepath.Join(corpusDir, fmt.Sprintf("%d", i))
		f, err := os.Create(fullPath)
		if err != nil {
			log.Fatalf("Failed to create corpus #%d: %q", i, corpus)
		}
		fmt.Fprintf(f, corpus)
		f.Close()
	}
	return nil
}
