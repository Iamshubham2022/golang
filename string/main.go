package main

import (
	"fmt"
	"strings"
)

func main() {
	// Original String
	str := "   Hello, Go Programming Language! Go is awesome.   "

	// 1. Trim leading and trailing whitespace
	trimmedStr := strings.TrimSpace(str)
	fmt.Println("Trimmed String:", trimmedStr) // Output: "Hello, Go Programming Language! Go is awesome."

	// 2. Split string into words
	words := strings.Split(trimmedStr, " ")
	fmt.Println("Words:", words) // Output: ["Hello," "Go" "Programming" "Language!" "Go" "is" "awesome."]

	// 3. Join words with a different separator
	joinedStr := strings.Join(words, "-")
	fmt.Println("Joined String:", joinedStr) // Output: "Hello,-Go-Programming-Language!-Go-is-awesome."

	// 4. Check if a substring exists
	containsGo := strings.Contains(trimmedStr, "Go")
	fmt.Println("Contains 'Go':", containsGo) // Output: true

	// 5. Replace all occurrences of a substring
	replacedStr := strings.ReplaceAll(trimmedStr, "Go", "Golang")
	fmt.Println("Replaced String:", replacedStr) // Output: "Hello, Golang Programming Language! Golang is awesome."

	// 6. Convert to lowercase
	lowercaseStr := strings.ToLower(trimmedStr)
	fmt.Println("Lowercase String:", lowercaseStr) // Output: "hello, go programming language! go is awesome."

	// 7. Convert to uppercase
	uppercaseStr := strings.ToUpper(trimmedStr)
	fmt.Println("Uppercase String:", uppercaseStr) // Output: "HELLO, GO PROGRAMMING LANGUAGE! GO IS AWESOME."

	// 8. Check prefix and suffix
	hasPrefix := strings.HasPrefix(trimmedStr, "Hello")
	hasSuffix := strings.HasSuffix(trimmedStr, "awesome.")
	fmt.Println("Has Prefix 'Hello':", hasPrefix)   // Output: true
	fmt.Println("Has Suffix 'awesome.':", hasSuffix) // Output: true

	// 9. Count occurrences of a substring
	countGo := strings.Count(trimmedStr, "Go")
	fmt.Println("Count of 'Go':", countGo) // Output: 2

	// 10. Repeat a string
	repeatedStr := strings.Repeat("Go ", 3)
	fmt.Println("Repeated String:", repeatedStr) // Output: "Go Go Go "

	// 11. Index of a substring
	index := strings.Index(trimmedStr, "Programming")
	fmt.Println("Index of 'Programming':", index) // Output: 10

	// 12. Last index of a substring
	lastIndex := strings.LastIndex(trimmedStr, "Go")
	fmt.Println("Last Index of 'Go':", lastIndex) // Output: 37

	// 13. Trim specific characters
	trimmedChars := strings.Trim(trimmedStr, " .!")
	fmt.Println("Trimmed Specific Characters:", trimmedChars) // Output: "Hello, Go Programming Language! Go is awesome"

	// 14. Check if string contains any characters from a set
	containsAny := strings.ContainsAny(trimmedStr, "xyz")
	fmt.Println("Contains any 'xyz':", containsAny) // Output: false

	// 15. Fields: Split by whitespace
	fields := strings.Fields(trimmedStr)
	fmt.Println("Fields:", fields) // Output: ["Hello," "Go" "Programming" "Language!" "Go" "is" "awesome."]

	// 16. Compare two strings (case-sensitive)
	compare := strings.Compare("abc", "ABC")
	fmt.Println("Compare 'abc' vs 'ABC':", compare) // Output: 1 (because "abc" > "ABC")

	// 17. EqualFold: Case-insensitive comparison
	equalFold := strings.EqualFold("abc", "ABC")
	fmt.Println("Case-insensitive comparison of 'abc' and 'ABC':", equalFold) // Output: true

	// 18. Title Case (First Letter Capitalized)
	titleCase := strings.Title("hello, go programming!")
	fmt.Println("Title Case String:", titleCase) // Output: "Hello, Go Programming!"

	// 19. Replace N occurrences of a substring
	replacedN := strings.Replace(trimmedStr, "Go", "Golang", 1)
	fmt.Println("Replace first 'Go' with 'Golang':", replacedN) // Output: "Hello, Golang Programming Language! Go is awesome."


	// we count the word how much time a word occur in the string 
}
