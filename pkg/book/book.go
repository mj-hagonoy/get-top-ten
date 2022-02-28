package book

import (
	"bufio"
	"regexp"
	"sort"
	"sync"
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
	contents        []byte
	wordCount       map[string]int
	wordCountMutext sync.RWMutex
}

func NewBook(input []byte) *Book {
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
	book.parseLines(book.contents)
}

func (book *Book) parseLines(buf []byte) {
	var wg sync.WaitGroup
	for {
		adv, token, err := bufio.ScanLines(buf, true)
		if adv == 0 || err != nil {
			break
		}
		wg.Add(1)
		go book.parseWords(&wg, token)
		if adv <= len(buf) {
			buf = buf[adv:]
		}
	}
	wg.Wait()
}

func (book *Book) parseWords(wg *sync.WaitGroup, buf []byte) {
	defer wg.Done()
	re := regexp.MustCompile("[a-zA-Z0-9]+")
	words := re.FindAllString(string(buf), -1)

	book.wordCountMutext.Lock()
	for _, word := range words {
		if ValidateWord(word) {
			book.wordCount[word]++
		}
	}
	book.wordCountMutext.Unlock()
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
func GetTopTenWords(contents []byte) []Rank {
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
