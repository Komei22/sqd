package detector

import (
	"bufio"
	"io"

	"github.com/Komei22/sqd/formatter"
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
	q := formatter.Format(query)
	q, err := parser.Parse(q)
	if err != nil {
		return "", err
	}
	if d.isSuspiciousQuery(q) {
		return query, nil
	}
	return "", nil
}

// DetectFrom query log file
func (d *Detector) DetectFrom(r io.Reader, suspiciousQueryChan chan<- string, errChan chan<- error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			errChan <- err
		}
		in := scanner.Text()
		query, err := formatter.ExtractQueryFrom(in)
		if err != nil {
			errChan <- err
		}
		suspiciousQuery, err := d.Detect(query)
		if err != nil {
			errChan <- err
		}
		if suspiciousQuery != "" {
			suspiciousQueryChan <- suspiciousQuery
		}
	}
	suspiciousQueryChan <- ""
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
