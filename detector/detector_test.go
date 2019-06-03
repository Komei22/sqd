package detector

import (
	"io"
	"strings"
	"testing"

	"github.com/Komei22/sqd/matcher"
	"github.com/Komei22/sql-mask"
)

func TestDetector_DetectFrom(t *testing.T) {
	blacklist := `DROP DATABASE test
DROP TABLE article`
	whitelist := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	queries := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT 10
DELETE FROM articles WHERE articles.id = 1
INSERT INTO articles (title, content, created_at, updated_at) VALUES ('hoge', 'hogehoge', '2018-11-01 09:53:37', '2018-11-01 09:53:37')
DROP DATABASE test
DROP TABLE article`
	illegalQueries := []string{
		"DROP DATABASE test",
		"DROP TABLE article",
	}

	bm := matcher.NewMatcher()
	bm.ReadList(strings.NewReader(blacklist))
	wm := matcher.NewMatcher()
	wm.ReadList(strings.NewReader(whitelist))

	mysqlMsk := &masker.MysqlMasker{}

	type fields struct {
		mode    Mode
		matcher *matcher.Matcher
		masker  masker.Masker
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Detect suspicius mysql queries using blacklist",
			fields: fields{
				mode:    Blacklist,
				matcher: bm,
				masker:  mysqlMsk,
			},
			args: args{
				r: strings.NewReader(queries),
			},
		},
		{
			name: "Detect suspicius mysql queries using whitelist",
			fields: fields{
				mode:    Whitelist,
				matcher: wm,
				masker:  mysqlMsk,
			},
			args: args{
				r: strings.NewReader(queries),
			},
		},
	}

	suspiciousQueryChan := make(chan string)
	errChan := make(chan error)
	defer close(suspiciousQueryChan)
	defer close(errChan)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Detector{
				mode:    tt.fields.mode,
				matcher: tt.fields.matcher,
				masker:  tt.fields.masker,
			}
			go d.DetectFrom(tt.args.r, suspiciousQueryChan, errChan)
			suspiciousQueries := []string{}
			for {
				query := <-suspiciousQueryChan
				if query == "" {
					break
				}
				suspiciousQueries = append(suspiciousQueries, query)
			}
			if len(suspiciousQueries) != len(illegalQueries) {
				t.Error("Faild detecting suspicious.(len(suspiciousQueries) != len(illegalQueries))")
			}
			for idx, sq := range suspiciousQueries {
				if illegalQueries[idx] != sq {
					t.Errorf("suspiciousQuery: %s, want: %s", sq, illegalQueries[idx])
				}
			}
		})
	}
}
