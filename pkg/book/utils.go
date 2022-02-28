package book

import "strings"

var SingleLetterWords = map[string]string{
	"a": "A",
	"i": "I",
}

// ValidateWord validates 'word'
//
// For word with len=1, only "A" & "I" are accepted
func ValidateWord(word string) bool {
	if len(word) == 1 {
		_, ok := SingleLetterWords[strings.ToLower(word)]
		return ok
	}
	return true
}
