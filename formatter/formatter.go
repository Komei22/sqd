package formatter

import (
	"regexp"
	"strings"
)

var multiSpaceRegexp = regexp.MustCompile(" {2,}")

// Format remove control charactor
func Format(query string) string {
	query = strings.Replace(query, "\t", " ", -1)
	query = multiSpaceRegexp.ReplaceAllString(query, " ")
	return query
}
