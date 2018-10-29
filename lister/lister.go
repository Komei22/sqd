package lister

import (
	"bufio"
	"io"
	"os"

	"github.com/Komei22/sql-mask"
)

// Create whitelist
func Create(r io.Reader, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		query := scanner.Text()
		queryStruct, err := parser.Parse(query)
		if err != nil {
			return err
		}
		file.Write(([]byte)(queryStruct + "\n"))
	}
	return nil
}
