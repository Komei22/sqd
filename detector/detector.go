package detector

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Komei22/sqd/matcher"
	"github.com/Komei22/sql-mask"
)

// Mode of detector
type Mode int

const (
	// Whitelist mode
	Whitelist = iota
	// Blacklist mode
	Blacklist
)

// Detector struct
type Detector struct {
	queries []string
	mode    Mode
}

// New detector
func New(i interface{}, mode Mode) (*Detector, error) {
	d := &Detector{}
	d.mode = mode
	var err error
	switch value := i.(type) {
	case string:
		err = d.readFile(value)
	case []string:
		d.queries = value
	default:
		err = fmt.Errorf("Parameter is unkown type. [Value type: %T]", i)
		return nil, err
	}

	return d, err
}

func (d *Detector) readFile(filepath string) error {
	r, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		d.queries = append(d.queries, scanner.Text())
	}
	return nil
}

// Detect find suspicious queries using matcher(whitelist or blacklist based)
func (d *Detector) Detect(m *matcher.Matcher) ([]string, error) {
	var suspiciousQueries []string
	for _, query := range d.queries {
		query, err := parser.Parse(query)
		if err != nil {
			return nil, err
		}
		if d.isSuspiciousQuery(query, m) {
			suspiciousQueries = append(suspiciousQueries, query)
		}
	}
	return suspiciousQueries, nil
}

func (d *Detector) isSuspiciousQuery(query string, m *matcher.Matcher) bool {
	if d.mode == Whitelist {
		if !m.IsMatchList(query) {
			return true
		}
	} else {
		if m.IsMatchList(query) {
			return true
		}
	}
	return false
}
