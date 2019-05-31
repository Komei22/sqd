package lister

import (
	"strings"
	"testing"

	"github.com/Komei22/sql-mask"
	"github.com/deckarep/golang-set"
)

func TestCreateUniqueQuerylist(t *testing.T) {
	queries := `SELECT * FROM user WHERE name = "test"
SELECT * FROM user WHERE name = "test"
SELECT * FROM user`
	expectList := mapset.NewSet("SELECT * FROM user WHERE name = ?", "SELECT * FROM user")

	m := &masker.MysqlMasker{}
	list, _ := Create(strings.NewReader(queries), m)

	if !expectList.Equal(list) {
		t.Errorf("Unexpected list: %s", list)
		t.Errorf("Expected list: %s", expectList)
	}
}
