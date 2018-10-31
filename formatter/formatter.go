package formatter

import (
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
	query = strings.Replace(query, "\\", "", -1)
	query = multiSpaceRegexp.ReplaceAllString(query, " ")
	return query
}
