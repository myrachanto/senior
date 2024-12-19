package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := `some magic of fmt package -what's under the hood?
hello everyone 🤗 , wow Christmas is now just around the corner
so today I am just going to highlight a few points on how fmt works under the hood
If a stringer method is attached to a type, and you intend to print it?
 the fmt package will prioritize that string method- like in the example below.
the fmt package can also take 0 arguments to well lots and still work, the sum function tries to imitate that!
and don't get me started on how it converts various data types to strings by default.
There is a lot more of what the fmt package is capable of!
Happy holidays everyone!
`
	words := strings.Fields(strings.ToLower(text))
	// Create a map to store word frequencies
	wordFrequency := make(map[string]int)
	// Count each word's occurrences
	for _, word := range words {
		// Clean up punctuation
		word = strings.Trim(word, ",.!?-'")
		wordFrequency[word]++
	}
	type output struct {
		key   string
		value int
	}
	results := []output{}
	for key, val := range wordFrequency {
		results = append(results, output{key, val})
	}
	//sort
	sort.Slice(results, func(i, j int) bool {
		return results[i].value > results[j].value
	})

	// Print the 5 popular words
	for _, out := range results[:5] {
		fmt.Printf("%s: %d\n", out.key, out.value)
	}
}
