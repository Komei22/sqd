package eventor

import (
	"fmt"
	"io"
)

// Dump suspicious queryies
func Dump(w io.Writer, suspiciousQueries []string) {
	for _, sq := range suspiciousQueries {
		fmt.Fprintln(w, sq)
	}
}
