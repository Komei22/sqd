package eventor

import (
	"fmt"
	"io"
)

// Print suspicious queryies
func Print(w io.Writer, suspiciousQueries []string) {
	for _, sq := range suspiciousQueries {
		fmt.Fprintln(w, sq)
	}
}
