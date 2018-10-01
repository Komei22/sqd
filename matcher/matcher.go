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

// New initalize Matcher
func New(filepath string) *Matcher {
	m := new(Matcher)
	m.whitelist = mapset.NewSet()
	m.initWhitelist(filepath)
	return m
}

// InitWhitelist initalize whitelist
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
	fmt.Println(m.whitelist)
}
