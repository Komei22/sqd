package eventor

import (
	"fmt"
)

// DumpStdout dump suspicious query for stdout
func DumpStdout(suspiciousQueries []string) {
	for _, sq := range suspiciousQueries {
		fmt.Println(sq)
	}
}
