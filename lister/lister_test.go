package lister

import (
	"github.com/deckarep/golang-set"
	"strings"
	"testing"
)

func TestCreateUniqueQuerylist(t *testing.T) {
	queries := `"SELECT * FROM user WHERE name = "test""
"SELECT * FROM user WHERE name = "test""
"SELECT * FROM user"`
	expectList := mapset.NewSet("SELECT * FROM user WHERE name = ?", "SELECT * FROM user")

	list, _ := Create(strings.NewReader(queries))

	if !expectList.Equal(list) {
		t.Errorf("Unexpected list: %s", list)
		t.Errorf("Expected list: %s", expectList)
	}
}
