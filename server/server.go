package server

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sql-mask"
)

// Server struct
type Server struct {
	detector *detector.Detector
}

// New server
func New(d *detector.Detector) *Server {
	s := &Server{detector: d}
	return s
}

// Start server
func (s *Server) Start() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		go s.handleDetection(scanner.Text())
	}
}

func (s *Server) handleDetection(querylog string) {
	parsedQuery, err := parser.Parse(querylog)
	if err != nil {
		fmt.Println(err)
	}
	suspiciousQuery, err := s.detector.Detect(parsedQuery)
	if err != nil {
		fmt.Println(err)
	}
	if suspiciousQuery != "" {
		fmt.Println(querylog)
	}
}
