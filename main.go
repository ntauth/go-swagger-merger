package main

import (
	"flag"

	"github.com/ntauth/go-swagger-merger/helpers"
	"golang.org/x/exp/slices"
)

func main() {
	merger := helpers.NewMerger()

	var outputFileName string

	flag.StringVar(&outputFileName, "o", "apis.swagger.json", "")
	flag.Parse()

	// Sort the files lexicographically in reverse so that the swagger annotations
	// artifact always comes last. This is required so that the merged file contains
	// the annotations info.
	files := flag.Args()
	slices.SortFunc(files, func(f1, f2 string) bool {
		return f1 > f2
	})

	for _, f := range files {
		err := merger.AddFile(f)
		if err != nil {
			panic(err)
		}
	}

	err := merger.Save(outputFileName)
	if err != nil {
		panic(err)
	}
}
