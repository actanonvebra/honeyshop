package helpers

import (
	"log"
	"net/url"
	"regexp"
	"strings"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func DetectSQLInjection(input string) bool {
	decodedInput := input
	if strings.Contains(input, "%") {
		decoded, err := url.QueryUnescape(input)
		if err != nil {
			decodedInput = decoded
		}
	}
	log.Printf("Original Input: %s | Decoded Input: %s", input, decodedInput)
	sqlInjectionPattern := `(?i)(\bselect\b|\bunion\b|\bupdate\b|\binsert\b|\bdelete\b|\bdrop\b|\bcreate\b|--|;|'|"|\*|=|\bfrom\b|\bwhere\b|\bor\b)`

	re := regexp.MustCompile(sqlInjectionPattern)
	return re.MatchString(decodedInput)
}
