package helpers

import "strings"

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func DetectSQLInjection(input string) bool {
	patterns := []string{"' OR '1'='1", "--", "DROP TABLE", "' OR 1=1"}
	for _, pattern := range patterns {
		if strings.Contains(strings.ToLower(input), pattern) {
			return true
		}
	}
	return false
}
