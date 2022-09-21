package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	s := fmt.Sprintf("%x", b)[:length]

	return s
}

func GenerateRandomSentence(numWords int, maxWordLen int) string {
	rand.Seed(time.Now().UnixNano())
	randomWords := make([]string, numWords)

	for i := range randomWords {
		randomWords[i] = GenerateRandomString(rand.Intn(maxWordLen) + 1)
	}

	randomParagraph := strings.Join(randomWords, " ")

	return randomParagraph
}
