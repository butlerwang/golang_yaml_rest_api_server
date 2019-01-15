package lib

import (
	"strings"
)

func StripSpace(str string) string {
	return strings.Join(strings.Fields(str), "_")
}
