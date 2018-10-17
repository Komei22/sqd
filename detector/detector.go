package detector

import (
	"bufio"
	"io"

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
	mode    Mode
	matcher *matcher.Matcher
}

// New detector
func New(m *matcher.Matcher, mode Mode) *Detector {
	d := &Detector{}
	d.mode = mode
	d.matcher = m
	return d
}

// Detect suspicious query
func (d *Detector) Detect(query string) (string, error) {
	q, err := parser.Parse(query)
	if err != nil {
		return "", err
	}
	if d.isSuspiciousQuery(q) {
		return query, nil
	}
	return "", nil
}

// DetectFrom query log file
func (d *Detector) DetectFrom(r io.Reader) ([]string, error) {
	var suspiciousQueries []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		query := scanner.Text()
		q, err := d.Detect(query)
		if err != nil {
			return nil, err
		}
		if q != "" {
			suspiciousQueries = append(suspiciousQueries, query)
		}
	}
	return suspiciousQueries, nil
}

func (d *Detector) isSuspiciousQuery(query string) bool {
	if d.mode == Whitelist {
		if !d.matcher.IsMatchList(query) {
			return true
		}
	} else {
		if d.matcher.IsMatchList(query) {
			return true
		}
	}
	return false
}
