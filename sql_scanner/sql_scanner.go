package sql_scanner

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/eventor"
	"github.com/Komei22/sql-mask"
)

// SqlScanner struct
type SqlScanner struct {
	detector *detector.Detector
}

// New SqlScanner
func New(d *detector.Detector) *SqlScanner {
	s := &SqlScanner{detector: d}
	return s
}

// Start sql_scanner
func (s *SqlScanner) Start() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		go s.handleDetection(scanner.Text())
	}
}

func (s *SqlScanner) handleDetection(querylog string) {
	parsedQuery, err := parser.Parse(querylog)
	if err != nil {
		fmt.Println(err)
	}
	suspiciousQuery, err := s.detector.Detect(parsedQuery)
	if err != nil {
		fmt.Println(err)
	}
	if suspiciousQuery != "" {
		eventor.DumpStdout([]string{suspiciousQuery})
	}
}
