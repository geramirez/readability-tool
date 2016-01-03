package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var PUNCTUATION = regexp.MustCompile("[.?!]+\\s")
var WORD = regexp.MustCompile("\\w+")

func countSentences(text string) int {
	return len(PUNCTUATION.Split(text, -1))
}

func getWords(text string) []string {
	return WORD.FindAllString(text, -1)
}

func getLetterType(letter rune) int {
	switch letter {
	case 'i', 'u':
		return 1
	case 'a', 'e', 'o', 'í', 'ú':
		return 2
	}
	return 3
}

func countSyllables(word string) int {
	word = strings.ToLower(word)
	var letterType, consonantRun, vowelRun, syllables int
	syllables = 1
	wordLength := utf8.RuneCountInString(word)
	for idx, letter := range word {
		letterType = getLetterType(letter)
		if letterType == 3 {
			if idx == 0 && wordLength > 3 {
				syllables--
			}
			consonantRun++
		} else {
			if consonantRun > 0 {
				syllables++
			}
			consonantRun = 0
		}
	}
	for _, letter := range word {
		letterType = getLetterType(letter)
		if letterType < 3 {
			vowelRun += letterType
		} else {
			if vowelRun > 3 {
				syllables++
			}
			vowelRun = 0
		}

	}
	if vowelRun > 3 && wordLength > 3 {
		syllables++
	}

	return syllables
}

func getCounts(text string) (int, int, int) {
	var totalSyllables, totalWords int
	for _, word := range getWords(text) {
		totalWords++
		totalSyllables += countSyllables(word)
	}
	return totalSyllables, totalWords, countSentences(text)
}

func main() {
	syllables, words, sentences := getCounts("how are you doing? and why are you doing it. how do you know?")
	fmt.Println(syllables, words, sentences)

}
