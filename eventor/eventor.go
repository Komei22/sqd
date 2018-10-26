package eventor

import (
	"fmt"
	"io"
)

// Print suspicious queryies
func Print(w io.Writer, suspiciousQueryChan <-chan string, errChan <-chan error) error {
	for {
		select {
		case suspiciousQuery := <-suspiciousQueryChan:
			if suspiciousQuery == "" {
				return nil
			}
			fmt.Fprintln(w, suspiciousQuery)
		case err := <-errChan:
			if err != nil {
				return err
			}
		}
	}
}
