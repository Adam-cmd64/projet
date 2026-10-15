package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Process(line string) string {
	line = strings.TrimSuffix(line, "\r")

	words := strings.Fields(line)
	words = joinMarkers(words)
	words = applyMarkers(words)
	words = fixPunctuation(words)
	words = fixQuotes(words)
	words = fixArticles(words)

	return strings.Join(words, " ")
}

func joinMarkers(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		w := words[i]
		if strings.HasPrefix(w, "(") && strings.HasSuffix(w, ",") && i+1 < len(words) {
			w = w + words[i+1]
			i++
		}
		result = append(result, w)
	}
	return result
}

func parseMarker(word string) (string, int, bool) {
	if !strings.HasPrefix(word, "(") || !strings.HasSuffix(word, ")") {
		return "", 0, false
	}
	inside := word[1 : len(word)-1]
	parts := strings.Split(inside, ",")

	cmd := parts[0]
	if cmd != "hex" && cmd != "bin" && cmd != "up" && cmd != "low" && cmd != "cap" {
		return "", 0, false
	}

	n := 1
	if len(parts) == 2 {
		number, err := strconv.Atoi(parts[1])
		if err != nil || number < 0 {
			return "", 0, false
		}
		n = number
	}
	if len(parts) > 2 {
		return "", 0, false
	}
	return cmd, n, true
}

func applyMarkers(words []string) []string {
	result := []string{}
	for _, w := range words {
		cmd, n, ok := parseMarker(w)
		if !ok {
			result = append(result, w)
			continue
		}
		for i := len(result) - 1; i >= 0 && n > 0; i-- {
			if !isWord(result[i]) {
				continue
			}
			result[i] = transform(cmd, result[i])
			n--
		}
	}
	return result
}

func transform(cmd string, word string) string {
	if cmd == "hex" {
		return toDecimal(word, 16)
	}
	if cmd == "bin" {
		return toDecimal(word, 2)
	}
	if cmd == "up" {
		return strings.ToUpper(word)
	}
	if cmd == "low" {
		return strings.ToLower(word)
	}
	if cmd == "cap" {
		return capitalize(word)
	}
	return word
}

func toDecimal(word string, base int) string {
	number, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(number))
}

func capitalize(word string) string {
	letters := []rune(strings.ToLower(word))
	if len(letters) == 0 {
		return word
	}
	letters[0] = unicode.ToUpper(letters[0])
	return string(letters)
}

func isWord(word string) bool {
	for _, c := range word {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func isPunct(c rune) bool {
	return strings.ContainsRune(".,!?:;", c)
}

func fixPunctuation(words []string) []string {
	result := []string{}
	for _, w := range words {
		count := 0
		for _, c := range w {
			if !isPunct(c) {
				break
			}
			count++
		}

		if count == 0 || len(result) == 0 {
			result = append(result, w)
			continue
		}

		punct := w[:count]
		rest := w[count:]

		result[len(result)-1] += punct
		if rest != "" {
			result = append(result, rest)
		}
	}
	return result
}

func fixQuotes(words []string) []string {
	result := []string{}
	open := false
	attachNext := false

	for _, w := range words {
		if w == "'" {
			if open && len(result) > 0 {
				result[len(result)-1] += "'"
				open = false
			} else {
				attachNext = true
				open = true
			}
			continue
		}
		if attachNext {
			w = "'" + w
			attachNext = false
		}
		result = append(result, w)
	}
	if attachNext {
		result = append(result, "'")
	}
	return result
}

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] != "a" && words[i] != "A" {
			continue
		}
		next := strings.TrimLeft(words[i+1], "'")
		if next == "" {
			continue
		}
		first := []rune(next)[0]
		if strings.ContainsRune("aeiouhAEIOUH", first) {
			words[i] = words[i] + "n"
		}
	}
	return words
}