package formatter

import (
	"testing"
)

func TestFormat(t *testing.T) {
	query := `"SELECT * FROM users\nWHERE\n\tname = "test""`
	expect := `SELECT * FROM users WHERE name = "test"`

	formattedQuery := Format(query)

	if formattedQuery != expect {
		t.Errorf("Unexpected fomatted query: %s", formattedQuery)
	}
}
