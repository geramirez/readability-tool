package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

//PUNCTUATION is a regex string for finding ".", "?", and "!".
var PUNCTUATION = regexp.MustCompile("[.?!]+\\s")

// WORD is a regex string for words
var WORD = regexp.MustCompile("\\w+")

func countSentences(text string) int {
	return len(PUNCTUATION.Split(text, -1))
}

// getWords
func getWords(text string) []string {
	return WORD.FindAllString(text, -1)
}

// getLetterType returns the type of letter
// 1 for weak vowel
// 2 for strong vowel
// 3 for consonant
func getLetterType(letter rune) int {
	switch letter {
	case 'i', 'u':
		return 1
	case 'a', 'e', 'o', 'í', 'ú':
		return 2
	}
	return 3
}

// countSyllables calculates the number of syllables using a modified
// version of the syllabification algorithm for spanish found in
// Heriberto Cuayáhuitl's article: http://www.dfki.de/~hecu01/publications/hc-cicling2004.pdf
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

// getCounts calculates syllables, words, and sentences
func getCounts(text string) (int, int, int) {
	var totalSyllables, totalWords int
	for _, word := range getWords(text) {
		totalWords++
		totalSyllables += countSyllables(word)
	}
	return totalSyllables, totalWords, countSentences(text)
}

// getStats calculates Fernandez Huerta's readability scores using an updated formula found here:
// http://linguistlist.org/issues/22/22-2332.html
// This function also returns syllables, words, and sentences.
func getStats(text string) (int, int, int, float32) {
	syllables, words, sentences := getCounts(text)
	readability := 206.84 - 60.0*(float32(syllables)/float32(words)) - 102.0*(float32(sentences)/float32(words))
	return syllables, words, sentences, readability
}

func main() {
	syllables, words, sentences, readability := getStats("Hello")
	fmt.Println(syllables, words, sentences, readability)

}
