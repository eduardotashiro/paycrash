package document

import (
	"regexp"
	"strings"
)

func removeCharacters(doc string) string {
	re := regexp.MustCompile(`[^A-Z0-9]+`)
	return re.ReplaceAllString(strings.ToUpper(doc), "")
}
