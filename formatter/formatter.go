package formatter

import (
	"fmt"
	"regexp"
	"strings"
)

var multiSpaceRegexp = regexp.MustCompile(" {2,}")

// Format remove control charactor
func Format(query string) string {
	removeStr := []string{"\\n", "\\t"}
	for _, str := range removeStr {
		query = strings.Replace(query, str, " ", -1)
	}
	query = strings.Replace(query, `\\`, `\`, -1)
	query = strings.Replace(query, `\"`, `"`, -1)
	query = strings.Replace(query, `\'`, `'`, -1)
	query = multiSpaceRegexp.ReplaceAllString(query, " ")
	return query
}

// ExtractQueryFrom input
func ExtractQueryFrom(in string) (string, error) {
	lastIdx := len(in) - 1
	if string(in[0]) != "\"" || string(in[lastIdx]) != "\"" {
		return "", fmt.Errorf("Invalid input format: %s", in)
	}
	return in[1:lastIdx], nil
}
