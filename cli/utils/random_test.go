package utils_test

import (
	"gitmsg/utils"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGenerateRandomString(t *testing.T) {
	length := 10
	randomStringLength := utf8.RuneCountInString(utils.GenerateRandomString(length))

	if randomStringLength != length {
		t.Errorf("utils.GenerateRandomString returned string of length %d, expected length %d", randomStringLength, length)
	}
}

func TestGenerateRandomSentence(t *testing.T) {
	numWords := 20
	maxWordLen := 5

	randomSentence := utils.GenerateRandomSentence(numWords, maxWordLen)
	words := strings.Fields(randomSentence)

	if len(words) != numWords {
		t.Errorf("utils.GenerateRandomSentence returned sentence with %d words, expected %d words", len(words), numWords)
	}

	for _, word := range words {
		wordLen := utf8.RuneCountInString(word)
		if wordLen > maxWordLen {
			t.Errorf("utils.GenerateRandomSentence returned sentence with a word of length %d, expected length %d or less", wordLen, maxWordLen)
		}
	}
}
