package detector

import (
	"strings"
	"testing"

	"github.com/Komei22/sqd/matcher"
)

func TestDetectSuspiciousQueriesUsingBlacklist(t *testing.T) {
	blacklist := `DROP DATABASE test
DROP TABLE article`

	queries := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)
DROP DATABASE test
DROP TABLE article`

	illegalQueries := []string{
		"DROP DATABASE test",
		"DROP TABLE article",
	}

	m := matcher.NewMatcher()
	m.ReadList(strings.NewReader(blacklist))

	d := New(m, Blacklist)
	suspiciousQueries, _ := d.DetectFrom(strings.NewReader(queries))

	for idx, sq := range suspiciousQueries {
		if illegalQueries[idx] != sq {
			t.Error("Failed detect suspicious queries based on blacklist.")
		}
	}
}

func TestDetectSuspiciousQueriesUsingwhitelist(t *testing.T) {
	whitelist := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`

	queries := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)
DROP DATABASE test
DROP TABLE article`

	illegalQueries := []string{
		"DROP DATABASE test",
		"DROP TABLE article",
	}

	m := matcher.NewMatcher()
	m.ReadList(strings.NewReader(whitelist))

	d := New(m, Whitelist)
	suspiciousQueries, _ := d.DetectFrom(strings.NewReader(queries))

	for idx, sq := range suspiciousQueries {
		if illegalQueries[idx] != sq {
			t.Error("Failed detect suspicious queries based on blacklist.")
		}
	}
}
