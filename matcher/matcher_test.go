package matcher

import (
	"io"
	"strings"
	"testing"
)

func TestMatcher_ReadList(t *testing.T) {
	type args struct {
		r io.Reader
	}

	queryList := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	// legalQueries := []string{
	// 	"SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?",
	// 	"DELETE FROM articles WHERE articles.id = ?",
	// 	"INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)",
	// }
	// illegalQuery := `DROP DATABASE production`

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Read queries and save mapset.Set",
			args:    args{r: strings.NewReader(queryList)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMatcher()
			if err := m.ReadList(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Matcher.ReadList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMatcher_IsMatchList(t *testing.T) {
	queryList := `SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	tests := []struct {
		name       string
		argQueries []string
		want       bool
	}{
		{
			name: "Is match list for listing query",
			argQueries: []string{
				"SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?",
				"DELETE FROM articles WHERE articles.id = ?",
				"INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)",
			},
			want: true,
		},
		{
			name: "Is match list for not listing query",
			argQueries: []string{
				"DROP DATABASE production",
				"DELETE FROM",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMatcher()
			m.ReadList(strings.NewReader(queryList))
			for _, arg := range tt.argQueries {
				if got := m.IsMatchList(arg); got != tt.want {
					t.Errorf("Matcher.IsMatchList() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
