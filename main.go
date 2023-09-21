package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	query := flag.String("query", "", "The filename to search for")
	flag.Parse()

	if len(*query) == 0 {
		fmt.Println("Error: No query provided")
		return
	}
	currPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error with getting working directory")
		return
	}
	fmt.Printf("Searching for %s...\n", *query)
	numResults := 0

	error := filepath.Walk(currPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(info.Name(), *query) {
			fmt.Printf("%s\t%s\n", info.Name(), path)
			numResults++
		}

		return nil
	})
	if error != nil {
		fmt.Println(error.Error())
	}
	if numResults == 0 {
		fmt.Printf("No results found for %s\n", *query)
	} else {
		fmt.Printf("Found %v results", numResults)
	}

}
