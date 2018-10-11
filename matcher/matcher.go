package matcher

import (
	"bufio"
	"io"
	"os"

	"github.com/deckarep/golang-set"
)

// Matcher struct
type Matcher struct {
	list mapset.Set
}

// NewMatcher ensure Matcher struct in memory
func NewMatcher() *Matcher {
	m := &Matcher{}
	m.list = mapset.NewSet()
	return m
}

// New initialize Matcher
func New(filepath string) (*Matcher, error) {
	m := NewMatcher()
	err := m.loadList(filepath)
	return m, err
}

func (m *Matcher) loadList(filepath string) error {
	reader, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer reader.Close()

	return m.SaveList(reader)
}

// SaveList save querys to list
func (m *Matcher) SaveList(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		m.list.Add(scanner.Text())
	}
	return nil
}

// IsMatchList returns true if the query is included in the list
func (m *Matcher) IsMatchList(query string) bool {
	return m.list.Contains(query)
}
