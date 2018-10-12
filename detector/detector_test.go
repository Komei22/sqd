package detector

import (
	"strings"
	"testing"

	"github.com/Komei22/sqd/matcher"
	"github.com/deckarep/golang-set"
)

func TestDetectSuspiciousQuerysUsingBlacklist(t *testing.T) {
	blacklist := `DROP DATABASE test
DROP TABLE article`

	querys := []string{
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

	d, _ := New(querys, Blacklist)
	suspiciousQuerys, _ := d.Detect(m)

	suspiciousQuerysSet := mapset.NewSet()
	for _, q := range suspiciousQuerys {
		suspiciousQuerysSet.Add(q)
	}
	if !suspiciousQuerysSet.Equal(illegalQuerySet) {
		t.Error("Failed detect suspicious querys based on blacklist.")
	}
}

func TestDetectSuspiciousQuerysUsingwhitelist(t *testing.T) {
	whitelist := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`

	querys := []string{
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

	d, _ := New(querys, Whitelist)
	suspiciousQuerys, _ := d.Detect(m)

	suspiciousQuerysSet := mapset.NewSet()
	for _, q := range suspiciousQuerys {
		suspiciousQuerysSet.Add(q)
	}
	if !suspiciousQuerysSet.Equal(illegalQuerySet) {
		t.Error("Failed detect suspicious querys based on whitelist.")
	}
}
