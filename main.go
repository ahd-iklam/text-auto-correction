package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Args := os.Args
	if len(Args) != 3 {
		fmt.Println("please enter a valid arguments <inputFile> <outputFile>")
		return
	}

	outputFile := os.Args[2]
	inputFile := os.Args[1]
	counter := 0

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Retrieve the input phrase from command-line arguments
	phrase := string(content)

	// Split the input phrase into words
	words := strings.Split(string(content), " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			if i == 0 {
				words[i] = ""
				continue
			}
			words[i-1] = strings.ToUpper(words[i-1]) // replace the previous word with uppercase letters
			words = removeElement(words, i)          // to remove the index that refer to the action needed
		} else if words[i] == "(low)" {
			if i == 0 {
				words[i] = ""
				continue
			}
			words[i-1] = strings.ToLower(words[i-1]) // replace the previous word with lowercase letters
			words = removeElement(words, i)
		} else if words[i] == "(cap)" {
			if i == 0 {
				words[i] = ""
				continue
			}
			words[i-1] = strings.Title(words[i-1]) // replace the first char of the previous letter with uppercase
			words = removeElement(words, i)
		} else if i+1 < len(words) && words[i] == "(up," {
			count, err := getCount(words[i+1])
			if err == nil && strings.Contains(words[i+1], ")") {
				v := i - count
				if v < 0 {
					v = 0
				}
				for j := v; j < i; j++ {
					words[j] = strings.ToUpper(words[j])
				}
				words = removeElement(words, i+1)
				words = removeElement(words, i)
			}

		} else if i+1 < len(words) && words[i] == "(low," {
			count, err := getCount(words[i+1])
			if err == nil && strings.Contains(words[i+1], ")") {
				v := i - count
				if v < 0 {
					v = 0
				}
				for j := v; j < i; j++ {
					words[j] = strings.ToLower(words[j])
				}
				words = removeElement(words, i+1)
				words = removeElement(words, i)
			}

		} else if i+1 < len(words) && words[i] == "(cap," {
			count, err := getCount(words[i+1])
			if err == nil && strings.Contains(words[i+1], ")") {
				v := i - count
				if v < 0 {
					v = 0
				}
				for j := v; j < i; j++ {
					words[j] = strings.Title(words[j])
				}
				words = removeElement(words, i+1)
				words = removeElement(words, i)
			}
		} else if words[i] == "(hex)" {
			if i == 0 {
				words[i] = ""
				continue
			}
			// Check if the preceding word is a hexadecimal number
			if val, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				// Convert the hexadecimal number to decimal
				words[i-1] = strconv.FormatInt(val, 10)
			}
			words = removeElement(words, i)
		} else if words[i] == "(bin)" {
			if i == 0 {
				words[i] = ""
				continue
			}
			// Check if the preceding word is a binary number
			if val, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				// Convert the binary number to decimal
				words[i-1] = strconv.FormatInt(val, 10)
			}
			words = removeElement(words, i)
			// If the word is "A" or "a" and the preceding word starts with a vowel or "h", replace it with "An" or "an"
		}
		if i+1 < len(words) && (words[i] == "A" || words[i] == "a") && isVowelOrH(string(words[i+1][0])) {
			if words[i] == "A" { // verify if the vowel letter is Lower/uppercase
				words[i] = "An"
			} else if words[i] == "a" {
				words[i] = "an"
			}
		} else if i+1 < len(words) && (words[i] == "An" || words[i] == "an") && !(isVowelOrH(string(words[i+1][0]))) {
			if words[i] == "An" { // verify if the vowel letter is Lower/uppercase
				words[i] = "A"
			} else if words[i] == "an" {
				words[i] = "a"
			}
		}
		if content[i] == '\'' {
			counter++
		}
		if counter%2 == 0 {
			iputedValue := regexp.MustCompile(`'\s*(.*?)\s*'`)
			phrase = iputedValue.ReplaceAllString(phrase, "'$1'")

		} else {
			iputedValue := regexp.MustCompile(`'\s*(.*?)\s*'\s*(.*?)\s*'`)
			phrase = iputedValue.ReplaceAllString(phrase, "'$1'$2'")
		}
	}
	// Call to the func to  Correct the phrase
	phrase = strings.Join(words, " ")
	phrase = Punctuation([]byte(phrase))
	phrase = handleSingleQuotes(phrase)

	err = ioutil.WriteFile(outputFile, []byte(phrase), 0o644) // Writes the corrected phrase to "result.txt"
	if err != nil {
		fmt.Println(err.Error()) // Prints an error message if writing fails
		return
	}
	//fmt.Println(phrase)
}

func isVowelOrH(char string) bool {
	vowels := "aeiouhAEIOUH"
	return strings.Contains(vowels, char)
}

func Punctuation(content []byte) string {
	ponctions := []byte{',', ';', ':', '!', '?', '.'}
	bol := true
	for bol {
		bol = false
		for i := 1; i <= len(content); i++ {
			for j := 0; j < len(ponctions); j++ {
				if i < len(content) && content[i-1] == ' ' && content[i] == ponctions[j] {
					content[i-1] = ponctions[j]
					content[i] = ' '
					bol = true
				}
			}
		}
	}
	re := regexp.MustCompile(` {2,}`)
	// Remplacer les deux espaces ou plus par un seul espace
	newText := re.ReplaceAllString(string(content), " ")
	return string(newText)
}

func removeElement(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func getCount(word string) (int, error) {
	word = strings.TrimPrefix(word, ",")
	word = strings.TrimSuffix(word, ")")
	return strconv.Atoi(word)
}

func isVowel(c byte) bool {
	vowels := "aeiouAEIOUhH"
	return strings.ContainsRune(vowels, rune(c))
}

func handleSingleQuotes(phrase string) string {
	count := 0
	for i := 0; i < len(phrase); i++ {
		if phrase[i] == '\'' {
			count++
		}
	}
	if count%2 == 0 {
		iputedValue := regexp.MustCompile(`'\s*(.*?)\s*'`)
		phrase := iputedValue.ReplaceAllString(phrase, "'$1'")
		return phrase

	} else {
		iputedValue := regexp.MustCompile(`'\s*(.*?)\s*'\s*(.*?)\s*'`)
		phrase := iputedValue.ReplaceAllString(phrase, "'$1'$2'")
		return phrase
	}
}
