package book

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestScanWords(t *testing.T) {
	testCase := `The greatest glory in living lies not in never falling, but in rising every time we fall.`
	expectedWords := []string{`The`, `greatest`, `glory`, `in`, `living`, `lies`, `not`, `never`, `falling`, `but`, `rising`, `every`, `time`, `we`, `fall`}

	start := time.Now()
	book := NewBook([]byte(testCase))
	book.ScanWords()
	elapsed := time.Since(start)
	t.Logf("book.ScanWords() completed in %s", elapsed)
	actualWords := book.GetWords()
	for _, word := range expectedWords {
		if _, ok := actualWords[word]; !ok {
			t.Fatalf("word '%s' not found", word)
		}
	}
}

func TestGetTopTenWord(t *testing.T) {
	testCases := []string{
		`.\sample\moby-dick.txt`,
		`.\sample\short-text.txt`,
		`.\sample\the-divine-comedy.txt`,
		`.\sample\the-king-james-bible.txt`,
	}

	for _, file := range testCases {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatalf("file %s cause error %s", file, err.Error())
		}
		start := time.Now()
		top10 := GetTopTenWords(data)
		elapsed := time.Since(start)
		t.Logf("GetTopTenWords() for file %s completed in %s", file, elapsed)
		t.Logf("file: %s\ntop 10:\n%+v", file, top10)
	}
}
