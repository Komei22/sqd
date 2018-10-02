package matcher

import (
	"bufio"
	"os"

	"github.com/deckarep/golang-set"
)

// Matcher struct
type Matcher struct {
	whitelist mapset.Set
}

// New initialize Matcher
func New(filepath string) (*Matcher, error) {
	m := new(Matcher)
	m.whitelist = mapset.NewSet()
	err := m.initWhitelist(filepath)
	return m, err
}

func (m *Matcher) initWhitelist(filepath string) error {
	fp, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		m.whitelist.Add(scanner.Text())
	}
	return nil
}

// IsLegitimate returns true if the query is included in the whitelist
func (m *Matcher) IsLegitimate(query string) bool {
	set := mapset.NewSet()
	set.Add(query)
	return set.IsSubset(m.whitelist)
}
