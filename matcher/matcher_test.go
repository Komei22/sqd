package matcher

import (
	"strings"
	"testing"
)

func TestIsMatchListforListingQuery(t *testing.T) {
	queryList := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`

	queries := []string{
		"SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?",
		"DELETE FROM articles WHERE articles.id = ?",
		"INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)",
	}

	m := NewMatcher()
	m.SaveList(strings.NewReader(queryList))

	for _, query := range queries {
		if !m.IsMatchList(query) {
			t.Error("Failed distinguish legitimate query.")
		}
	}
}

func TestIsMatchListforNotListingQuery(t *testing.T) {
	queryList := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	query := "DROP DATABASE production"

	m := NewMatcher()
	m.SaveList(strings.NewReader(queryList))

	if m.IsMatchList(query) {
		t.Error("Failed distinguish illegal query.")
	}
}
