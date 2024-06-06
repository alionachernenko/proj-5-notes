package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filesDirectory := flag.String("dir", "./notes", "files directory")

	words := map[string][]string{
		"expression": {"functions.txt"},
		"variable":   {"functions.txt"},
		"sequence":   {"algorithms.txt"},
	}

	var searchInput string

	fmt.Scan(&searchInput)

	files := getFiles(filesDirectory, words[strings.ToLower(searchInput)])

	if len(files) == 0 {
		fmt.Println("Nothing was found :(")
		return
	}

	fmt.Println("Results:")

	for _, file := range files {
		fmt.Println(file)
	}
}

func getFiles(dir *string, fileNames []string) []string {
	var files []string

	for _, fileName := range fileNames {
		file, err := os.Open(fmt.Sprintf("%v/%v", *dir, fileName))

		if err != nil {
			fmt.Printf("Error opening file %v: %v", fileName, err)
			return nil
		}

		defer file.Close()
		
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			files = append(files, line)
		}
	}
	
	return files
}
