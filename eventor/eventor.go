package eventor

import (
	"fmt"
	"io"
	"os"
)

// Print suspicious queryies
func Print(w io.Writer, suspiciousQueryChan <-chan string, errChan <-chan error) {
	for {
		select {
		case suspiciousQuery := <-suspiciousQueryChan:
			if suspiciousQuery == "" {
				return
			}
			fmt.Fprintln(w, suspiciousQuery)
		case err := <-errChan:
			fmt.Fprintf(os.Stderr, "Can't detection suspicious query: %s", err)
		}
	}
}
