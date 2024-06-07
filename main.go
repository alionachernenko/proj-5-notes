package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func createWordsMap(fileContent string) (map[string][]int, []string) {
	words := make(map[string][]int)

	lines := strings.Split(strings.ToLower(fileContent), ".")

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		lineWords := strings.Fields(trimmedLine)

		for _, word := range lineWords {
			words[word] = append(words[word], i)
		}
	}

	return words, lines
}

func getFile(dir *string) (*os.File, error) {
	filePath := fmt.Sprintf("%v/note.txt", *dir)
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func getFileContent(file *os.File) string {
	var fileContent string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileContent = scanner.Text()
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

	words, lines := createWordsMap(fileContent)

	var input string

	fmt.Print("Enter word to search: ")

	fmt.Scan(&input)

	results := words[strings.ToLower(input)]

	if len(results) == 0 {
		fmt.Print("Nothing was found :(")
		return
	}

	for _, lineIdx := range results {
		fmt.Println(lines[lineIdx])
	}
}
