package formatter

import (
	"strings"
)

// TrimControlChara remove control charactor
func TrimControlChara(query string) string {
	query = query[1:strings.LastIndex(query, "\"")]
	removeStr := []string{"\\n", "\\t"}
	for _, str := range removeStr {
		query = strings.Replace(query, str, " ", -1)
	}
	query = strings.Replace(query, "\\", "", -1)
	return query
}
