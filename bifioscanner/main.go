package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func readByline(filename string) (map[string]int, error) {
	wordsMap := make(map[string]int)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//get each line on the file
		line := scanner.Text()
		// gets words
		words := strings.Fields(line)
		for _, word := range words {
			wordsMap[word]++
		}
	}
	// check for errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return wordsMap, nil
}
func sortMap(data map[string]int) []wordCount {
	// move data frm map to string
	wordCountSlice := make([]wordCount, 0, len(data))

	for k, v := range data {
		wordCountSlice = append(wordCountSlice, wordCount{k, v})
	}
	sort.Slice(wordCountSlice, func(i, j int) bool {
		return wordCountSlice[i].count > wordCountSlice[j].count
	})
	return wordCountSlice
}
func main() {
	filename := "./error-shadowing.txt"
	results, err := readByChunks(filename)
	if err != nil {
		log.Println(err)
	}
	// get top 10 most used words
	wordCount := sortMap(results)
	for _, v := range wordCount[0:10] {
		fmt.Printf("%s occured %d times thought the file \n", v.word, v.count)
	}
}
func readByChunks(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := io.Reader(file)
	bufferSize := 1024

	wordsMap := make(map[string]int)
	for {
		buffer := make([]byte, bufferSize)
		n, err := reader.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				return nil, fmt.Errorf("error %w", err)
			}
			break
		}
		words := strings.Fields(string(buffer[:n]))
		for _, word := range words {
			wordsMap[word]++
		}
	}
	return wordsMap, nil
}
