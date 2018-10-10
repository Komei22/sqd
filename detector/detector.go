package detector

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Komei22/sqd/matcher"
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

// DumpSuspiciousQuerys find suspicious querys using matcher(whitelist or blacklist based)
func (d *Detector) DumpSuspiciousQuerys(m *matcher.Matcher) {
	fmt.Printf("Suspicious querys:\n")
	for _, query := range d.querys {
		if d.isSuspiciousQuery(query, m) {
			fmt.Printf("%s\n", query)
		}
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
