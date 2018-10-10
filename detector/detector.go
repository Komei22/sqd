package detector

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Komei22/sqd/matcher"
	"github.com/Komei22/sql-mask"
)

// Detector struct
type Detector struct {
	querys []string
	mode   string
}

// New detector
func New(filepath string, mode string) (*Detector, error) {
	d := &Detector{}
	err := d.readQuerys(filepath)
	d.mode = mode
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

// Dump output suspiciousQuerys
func Dump(suspiciousQuerys []string) {
	fmt.Print("Suspicious querys\n")
	for _, query := range suspiciousQuerys {
		fmt.Printf("%s\n", query)
	}
}

func (d *Detector) isSuspiciousQuery(query string, m *matcher.Matcher) bool {
	if d.mode == "whitelist" {
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
