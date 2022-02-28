package book

import (
	"bufio"
	"sort"
	"strings"
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

func (book *Book) GetWords() map[string]int {
	return book.wordCount
}

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
	for {
		adv, token, err := bufio.ScanWords(buf, true)
		if adv == 0 || err != nil {
			break
		}
		word := cleanWord(string(token))

		book.wordCountMutext.Lock()
		book.wordCount[word]++
		book.wordCountMutext.Unlock()

		if adv <= len(buf) {
			buf = buf[adv:]
		}
	}
}
func cleanWord(word string) string {
	//word = strings.ToLower(word)
	word = strings.ReplaceAll(word, `.`, ``)
	word = strings.ReplaceAll(word, `'s`, ``)
	word = strings.ReplaceAll(word, `?`, ``)
	word = strings.ReplaceAll(word, `!`, ``)
	word = strings.ReplaceAll(word, `:`, ``)
	word = strings.ReplaceAll(word, `;`, ``)
	word = strings.ReplaceAll(word, `,`, ``)
	word = strings.ReplaceAll(word, `'`, ``)
	word = strings.ReplaceAll(word, `)`, ``)
	return word
}

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
