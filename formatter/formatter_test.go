package formatter

import (
	"testing"
)

func TestFormat(t *testing.T) {
	queries := []string{
		`SELECT * FROM         users\nWHERE\n\tname = \"test\"`,
		`SELECT * FROM users WHERE name = \'te\\\"st\'`}
	expects := []string{
		`SELECT * FROM users WHERE name = "test"`,
		`SELECT * FROM users WHERE name = 'te\"st'`}

	var formattedQueries []string
	for _, q := range queries {
		formattedQueries = append(formattedQueries, Format(q))
	}
	for i, fq := range formattedQueries {
		if fq != expects[i] {
			t.Errorf("Unexpected fomatted query: %s", fq)
		}
	}
}
