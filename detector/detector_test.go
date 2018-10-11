package detector

import (
	"strings"
	"testing"

	"github.com/Komei22/sqd/matcher"
	"github.com/deckarep/golang-set"
)

func TestDumpSuspiciousQuerysUsingBlacklist(t *testing.T) {
	blacklist := `DROP DATABASE test
DROP TABLE article`

	querys := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)
DROP DATABASE test
DROP TABLE article`

	illegalQuerySet := mapset.NewSet(
		"DROP DATABASE test",
		"DROP TABLE article",
	)

	m := matcher.NewMatcher()
	m.SaveList(strings.NewReader(blacklist))

	d := newDetector(Blacklist)
	d.saveQuerys(strings.NewReader(querys))

	suspiciousQuerys, _ := d.Detect(m)

	suspiciousQuerysSet := mapset.NewSet()
	for _, q := range suspiciousQuerys {
		suspiciousQuerysSet.Add(q)
	}
	if !suspiciousQuerysSet.Equal(illegalQuerySet) {
		t.Error("Failed detect suspicious querys based on blacklist.")
	}
}

func TestDumpSuspiciousQuerysUsingwhitelist(t *testing.T) {
	whitelist := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
	DELETE FROM articles WHERE articles.id = ?
	INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`

	querys := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)
DROP DATABASE test
DROP TABLE article`

	illegalQuerySet := mapset.NewSet(
		"DROP DATABASE test",
		"DROP TABLE article",
	)

	m := matcher.NewMatcher()
	m.SaveList(strings.NewReader(whitelist))

	d := newDetector(Whitelist)
	d.saveQuerys(strings.NewReader(querys))

	suspiciousQuerys, _ := d.Detect(m)

	suspiciousQuerysSet := mapset.NewSet()
	for _, q := range suspiciousQuerys {
		suspiciousQuerysSet.Add(q)
	}
	if suspiciousQuerysSet.Equal(illegalQuerySet) {
		t.Error("Failed detect suspicious querys based on whitelist.")
	}
}