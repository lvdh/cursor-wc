package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wc <file>")
		return
	}

	fileName := os.Args[1]
	count, err := countWords(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Word count: %d\n", count)
}

func countWords(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		wordCount += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}
