package sqlscanner

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/eventor"
	"github.com/Komei22/sql-mask"
)

// SQLScanner struct
type SQLScanner struct {
	detector *detector.Detector
}

// New SQLScanner
func New(d *detector.Detector) *SQLScanner {
	s := &SQLScanner{detector: d}
	return s
}

// Start sql_scanner
func (s *SQLScanner) Start(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		s.detection(scanner.Text())
	}
}

func (s *SQLScanner) detection(querylog string) {
	parsedQuery, err := parser.Parse(querylog)
	if err != nil {
		fmt.Println(err)
	}
	suspiciousQuery, err := s.detector.Detect(parsedQuery)
	if err != nil {
		fmt.Println(err)
	}
	if suspiciousQuery != "" {
		eventor.Print(os.Stdout, []string{suspiciousQuery})
	}
}
