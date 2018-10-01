package matcher

import (
	"bufio"
	"fmt"
	"os"
)

// Matcher struct
type Matcher struct {
	whitelist []string
}

// New initalize Matcher
func New(filepath string) *Matcher {
	m := new(Matcher)
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
		m.whitelist = append(m.whitelist, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
