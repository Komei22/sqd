package eventor

import (
	"fmt"
	"io"
)

// Print suspicious queryies
func Print(w io.Writer, c <-chan string) {
	for {
		q, ok := <-c
		if !ok {
			break
		}
		if q != "" {
			fmt.Fprintln(w, q)
		}
	}
}
