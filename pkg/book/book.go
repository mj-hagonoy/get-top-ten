package book

import (
	"regexp"
	"sort"
)

type Rank struct {
	Rank  int           `json:"rank"`
	Value WordFrequency `json:"value"`
}

type WordFrequency struct {
	Frequency int      `json:"frequency"`
	Words     []string `json:"words"`
}

type MapWordFrequency map[int]WordFrequency

type Book struct {
	contents  string
	wordCount map[string]int
}

func NewBook(input string) *Book {
	book := &Book{
		contents:  input,
		wordCount: make(map[string]int),
	}
	return book
}

// GetWords returns map of words found and corresponding count
//
// May return empty or nil when Book.ScanWords is not executed prior to GetWords
func (book *Book) GetWords() map[string]int {
	return book.wordCount
}

// ScanWords extracts words from Book.contents
func (book *Book) ScanWords() {
	book.parseWords(book.contents)
}

func (book *Book) parseWords(input string) {
	re := regexp.MustCompile("[a-zA-Z0-9]+")
	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		word := string(match)
		if ValidateWord(word) {
			book.wordCount[word]++
		}
	}
}

// groupWordsByFrequency accepts wordCount map
//
// Returns:
//
// result MapWordFrequency : map with list words grouped by frequency
//
// keys []int : keys in decreasing order
func groupWordsByFrequency(wordCount map[string]int) (result MapWordFrequency, keys []int) {
	result = make(MapWordFrequency)
	for word, frequency := range wordCount {
		element, ok := result[frequency]
		if ok {
			element.Words = append(element.Words, word)
			result[frequency] = element
		} else {
			result[frequency] = WordFrequency{
				Frequency: frequency,
				Words:     []string{word},
			}
			keys = append(keys, frequency)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	return result, keys
}

// GetTopTenWords accepts array bytes (text)
// and returns the top 10 most used words
func GetTopTenWords(contents string) []Rank {
	if len(contents) == 0 {
		return nil
	}
	book := NewBook(contents)
	book.ScanWords()

	groups, keys := groupWordsByFrequency(book.wordCount)

	size := 10
	if len(keys) < 10 {
		size = len(keys)
	}
	keys = keys[0:size]

	top10 := make([]Rank, 0)
	rank := 1
	for _, key := range keys {
		top10 = append(top10, Rank{
			Rank:  rank,
			Value: groups[key],
		})
		rank++
	}

	return top10
}
