package eventor

import (
	"fmt"
)

// DumpStdout dump suspicious query for stdout
func DumpStdout(suspiciousQueries []string) {
	fmt.Println("Suspicious query:")
	for _, sq := range suspiciousQueries {
		fmt.Println(sq)
	}
}
