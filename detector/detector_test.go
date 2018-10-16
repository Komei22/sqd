package detector

import (
	"strings"
	"testing"

	"github.com/Komei22/sqd/matcher"
	"github.com/deckarep/golang-set"
)

func TestDetectSuspiciousQueriesUsingBlacklist(t *testing.T) {
	blacklist := `DROP DATABASE test
DROP TABLE article`

	queries := []string{
		"SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?",
		"DELETE FROM articles WHERE articles.id = ?",
		"INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)",
		"DROP DATABASE test",
		"DROP TABLE article",
	}

	illegalQuerySet := mapset.NewSet(
		"DROP DATABASE test",
		"DROP TABLE article",
	)

	m := matcher.NewMatcher()
	m.ReadList(strings.NewReader(blacklist))

	d := New(m, Blacklist)
	suspiciousQuerySet := mapset.NewSet()
	for _, q := range queries {
		sq, _ := d.Detect(q)
		if sq != "" {
			suspiciousQuerySet.Add(sq)
		}
	}

	if !suspiciousQuerySet.Equal(illegalQuerySet) {
		t.Error("Failed detect suspicious queries based on blacklist.")
	}
}

func TestDetectSuspiciousQueriesUsingwhitelist(t *testing.T) {
	whitelist := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`

	queries := []string{
		"SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?",
		"DELETE FROM articles WHERE articles.id = ?",
		"INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)",
		"DROP DATABASE test",
		"DROP TABLE article",
	}

	illegalQuerySet := mapset.NewSet(
		"DROP DATABASE test",
		"DROP TABLE article",
	)

	m := matcher.NewMatcher()
	m.ReadList(strings.NewReader(whitelist))

	d := New(m, Whitelist)
	suspiciousQuerySet := mapset.NewSet()
	for _, q := range queries {
		sq, _ := d.Detect(q)
		if sq != "" {
			suspiciousQuerySet.Add(sq)
		}
	}

	if !suspiciousQuerySet.Equal(illegalQuerySet) {
		t.Error("Failed detect suspicious queries based on whitelist.")
	}
}
