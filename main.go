package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func createWordsMap(fileContent []string) map[string][]int {
	words := make(map[string][]int)

	for i, line := range fileContent {
		normalizedLine := strings.ToLower(line)
		trimmedLine := strings.TrimSpace(normalizedLine)

		lineWords := strings.Fields(trimmedLine)

		for _, word := range lineWords {
			if slices.Contains(words[word], i) {
				continue
			}

			words[word] = append(words[word], i)
		}
	}

	fmt.Print(words)
	return words
}

func getFile(dir *string) (*os.File, error) {
	filePath := fmt.Sprintf("%v/notes.txt", *dir)
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func getFileContent(file *os.File) []string {
	var fileContent []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines := scanner.Text()
		fileContent = append(fileContent, lines)
	}

	return fileContent
}


func main() {
	filesDirectory := flag.String("dir", "./notes", "files directory")
	flag.Parse()

	file, err := getFile(filesDirectory)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	defer file.Close()

	fileContent := getFileContent(file)

	words := createWordsMap(fileContent)

	var input string

	fmt.Print("Enter word to search: ")

	fmt.Scan(&input)

	results := words[strings.ToLower(input)]
	fmt.Print(results)

	if len(results) == 0 {
		fmt.Print("Nothing was found :(")
		return
	}

	for _, lineIdx := range results {
		fmt.Println(fileContent[lineIdx])
	}
}
