package matcher

import (
	"bufio"
	"fmt"
	"os"

	"github.com/deckarep/golang-set"
)

// Matcher struct
type Matcher struct {
	whitelist mapset.Set
}

// New initialize Matcher
func New(filepath string) *Matcher {
	m := new(Matcher)
	m.whitelist = mapset.NewSet()
	m.initWhitelist(filepath)
	return m
}

func (m *Matcher) initWhitelist(filepath string) {
	fp, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		m.whitelist.Add(scanner.Text())
	}
}

// IsLegitimate returns true if the query is included in the whitelist
func (m *Matcher) IsLegitimate(query string) bool {
	set := mapset.NewSet()
	set.Add(query)
	return set.IsSubset(m.whitelist)
}
