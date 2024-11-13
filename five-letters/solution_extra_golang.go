package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Solver struct {
	words     []string
	available []rune
	solutions []string
}

func NewSolver(letters []rune, words []string) *Solver {
	return &Solver{
		words:     words,
		available: letters,
	}
}

func (s *Solver) Solve() {
	s.solveRecursive([]string{}, s.available, s.words)
}

func (s *Solver) solveRecursive(used []string, letters []rune, words []string) {
	if len(letters) == 0 {
		var solution string
		for _, word := range used {
			if len(word) == 5 {
				solution += word + " "
			}
		}
		s.solutions = append(s.solutions, strings.TrimSpace(solution))
		return
	}

	joinedWords := strings.Join(words, "")
	rarest := letters[0]
	for _, letter := range letters {
		if strings.Count(joinedWords, string(letter)) < strings.Count(joinedWords, string(rarest)) {
			rarest = letter
		}
	}

	for _, word := range append(words, strings.Repeat(string(rarest), len(letters)%5)) {
		if strings.ContainsRune(word, rarest) {
			newLetters := s.filterLetters(letters, word)
			newWords := s.filterWords(words, word)
			s.solveRecursive(append(used, word), newLetters, newWords)
		}
	}
}

func (s *Solver) filterLetters(letters []rune, word string) []rune {
	var result []rune
	for _, letter := range letters {
		if !strings.ContainsRune(word, letter) {
			result = append(result, letter)
		}
	}
	return result
}

func (s *Solver) filterWords(words []string, word string) []string {
	var result []string
	for _, w := range words {
		if !hasCommonLetters(w, word) {
			result = append(result, w)
		}
	}
	return result
}

func hasCommonLetters(word1, word2 string) bool {
	for _, l := range word1 {
		if strings.ContainsRune(word2, l) {
			return true
		}
	}
	return false
}

func loadWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordsMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 5 && uniqueLetters(word) {
			sortedWord := sortString(word)
			wordsMap[sortedWord] = word
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var words []string
	for _, word := range wordsMap {
		words = append(words, word)
	}
	return words, nil
}

func uniqueLetters(word string) bool {
	letterSet := make(map[rune]struct{})
	for _, letter := range word {
		if _, exists := letterSet[letter]; exists {
			return false
		}
		letterSet[letter] = struct{}{}
	}
	return true
}

func sortString(word string) string {
	letters := []rune(word)
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}

func main() {
	words, err := loadWords("words_alpha.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	solver := NewSolver([]rune("abcdefghijklmnopqrstuvwxyz"), words)
	solver.Solve()

	fmt.Println(len(solver.solutions))
}
