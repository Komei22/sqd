package detector

import (
	"bufio"
	"io"
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
	querys []string
	mode   Mode
}

func newDetector(mode Mode) *Detector {
	d := &Detector{}
	d.mode = mode
	return d
}

// New detector
func New(filepath string, mode Mode) (*Detector, error) {
	d := newDetector(mode)
	err := d.readQuerys(filepath)
	return d, err
}

func (d *Detector) readQuerys(filepath string) error {
	r, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer r.Close()

	return d.saveQuerys(r)
}

func (d *Detector) saveQuerys(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		d.querys = append(d.querys, scanner.Text())
	}
	return nil
}

// Detect find suspicious querys using matcher(whitelist or blacklist based)
func (d *Detector) Detect(m *matcher.Matcher) ([]string, error) {
	var suspiciousQuerys []string
	for _, query := range d.querys {
		query, err := parser.Parse(query)
		if err != nil {
			return nil, err
		}
		if d.isSuspiciousQuery(query, m) {
			suspiciousQuerys = append(suspiciousQuerys, query)
		}
	}
	return suspiciousQuerys, nil
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
