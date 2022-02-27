package book

import (
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
