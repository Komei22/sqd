package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Komei22/sqd/detector"
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

func (s *Server) handleDetection(querylogJSON string) {
	jsonBytes := ([]byte)(querylogJSON)
	var querylog map[string]interface{}
	json.Unmarshal(jsonBytes, &querylog)
	suspiciousQuery, err := s.detector.Detect(querylog["query"].(string))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(suspiciousQuery)
}
