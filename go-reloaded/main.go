package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		return
	}
	myArgs := os.Args[1:]
	content, err := os.ReadFile(myArgs[0])
	if err != nil {
		fmt.Println("Error file", err)
		os.Exit(1)
	}
	s := strings.Fields(string(content))

	for i, word := range s {
		if word == "(hex)" {
			data, _ := strconv.ParseInt(s[i-1], 16, 64)
			s[i-1] = fmt.Sprint(data)
			s = append(s[:i], s[i+1:]...)
		}
	}

	for i, word := range s {
		if word == "(bin)" {
			data, _ := strconv.ParseInt(s[i-1], 2, 64)
			s[i-1] = fmt.Sprint(data)
			s = append(s[:i], s[i+1:]...)
		}
	}

	// modify out input text
	var1 := atoAn(s)
	var2 := lowcase(var1)
	var3 := UppCase(var2)
	var4 := capiTaliser(var3)
	var5 := punctuation(var4)
	var6 := strings.Join(var5, " ")
	var7 := []byte(var6)
	os.WriteFile(myArgs[1], var7, 0o644)
}

func lowcase(s []string) []string {

	for i, word := range s {
		if strings.Contains(word, "(low") {
			if strings.Contains(word, "(low,") {
				val, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for y := i - val; y < i; y++ {
					s[y] = strings.ToLower(s[y])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToLower(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

func UppCase(s []string) []string {
	for i, word := range s {
		if strings.Contains(word, "(up") {
			if strings.Contains(word, "(up,") {
				val, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for y := i - val; y < i; y++ {
					s[y] = strings.ToUpper(s[y])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToUpper(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

func capiTaliser(s []string) []string {
	for i, word := range s {
		if strings.Contains(word, "(cap") {
			if strings.Contains(word, "(cap,") {
				val, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for y := i - val; y < i; y++ {
					s[y] = strings.Title(s[y])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.Title(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

func atoAn(s []string) []string {
	vowelSample := []byte{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}
	for i, word := range s {
		for _, vowel := range vowelSample {
			if word == "a" && s[i+1][0] == vowel {
				s[i] = "an"
			} else if word == "A" && s[i+1][0] == vowel {
				s[i] = "An"
			}
		}
	}
	return s
}

func punctuation(s []string) []string {
	puns := []string{",", ".", "!", "?", ":", ";"}
	for i, word := range s {
		for _, puncs := range puns {
			if (string(word[0]) == puncs) && string(word[len(word)-1]) != puncs {
				s[i-1] = s[i-1] + puncs
				s[i] = word[1:]
			}
		}
	}

	for i, word := range s {
		for _, puncs := range puns {
			// the word begin with a punctuation and is the last word in the string
			if (string(word[0]) == puncs) && (s[i] == s[len(s)-1] ) {
				s[i-1] = s[i-1] + word
				s = s[:len(s)-1]
			}
		}
	}

	for i, word := range s {
		for _, puncs := range puns {
			if string(word[0]) == puncs && string(word[len(word)-1]) == puncs && s[i] != s[len(s)-1] {
				s[i-1] = s[i-1] + word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count = count + 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}

	for i, word := range s {
		if word == "'" && count != 0 {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}
