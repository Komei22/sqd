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
	m.SaveList(strings.NewReader(blacklist))

	d, _ := New(queries, Blacklist)
	suspiciousQueries, _ := d.Detect(m)

	suspiciousQueriesSet := mapset.NewSet()
	for _, q := range suspiciousQueries {
		suspiciousQueriesSet.Add(q)
	}
	if !suspiciousQueriesSet.Equal(illegalQuerySet) {
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
	m.SaveList(strings.NewReader(whitelist))

	d, _ := New(queries, Whitelist)
	suspiciousQueries, _ := d.Detect(m)

	suspiciousQueriesSet := mapset.NewSet()
	for _, q := range suspiciousQueries {
		suspiciousQueriesSet.Add(q)
	}
	if !suspiciousQueriesSet.Equal(illegalQuerySet) {
		t.Error("Failed detect suspicious queries based on whitelist.")
	}
}
