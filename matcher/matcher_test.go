package matcher

import (
	"github.com/deckarep/golang-set"
	"strings"
	"testing"
)

func TestDistinguishLegitimateQuery(t *testing.T) {
	queryList := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`

	querys := []string{
		"SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?",
		"DELETE FROM articles WHERE articles.id = ?",
		"INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)",
	}

	m := new(Matcher)
	m.list = mapset.NewSet()
	m.saveList(strings.NewReader(queryList))

	for _, query := range querys {
		if !m.IsMatchList(query) {
			t.Error("Failed distinguish legitimate query.")
		}
	}
}

func TestDistinguishIllegalQuery(t *testing.T) {
	queryList := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	query := "DROP DATABASE production"

	m := new(Matcher)
	m.list = mapset.NewSet()
	m.saveList(strings.NewReader(queryList))

	if m.IsMatchList(query) {
		t.Error("Failed distinguish illegal query.")
	}
}
