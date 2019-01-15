package lib

import (
	"strings"
)

func StripSpace(str string) string {
	return strings.Join(strings.Fields(str), "_")
}

func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}
