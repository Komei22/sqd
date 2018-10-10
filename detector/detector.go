package detector

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Komei22/sqd/matcher"
)

// DetectionMode is setted whitelist or blacklist mode
type DetectionMode int

const (
	// Whitelist mode detection
	Whitelist = iota
	// Blacklist mode detection
	Blacklist
)

// Detector struct
type Detector struct {
	querys []string
	mode   DetectionMode
}

// New detector
func New(filepath string, mode DetectionMode) (*Detector, error) {
	d := &Detector{}
	err := d.readQuerys(filepath)
	d.mode = mode
	return d, err
}

func (d *Detector) readQuerys(filepath string) error {
	reader, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
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
	var suspiciousQuerys []string
	for _, query := range d.querys {
		if d.isSuspiciousQuery(query, m) {
			suspiciousQuerys = append(suspiciousQuerys, query)
		}
	}
	fmt.Printf("Suspicious querys:\n")
	fmt.Print(suspiciousQuerys)
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
