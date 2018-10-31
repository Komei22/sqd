package lister

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Komei22/sqd/formatter"
	"github.com/Komei22/sql-mask"
	"github.com/deckarep/golang-set"
)

// Create whitelist
func Create(r io.Reader, w io.Writer) error {
	list := mapset.NewSet()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		query := formatter.Format(scanner.Text())
		queryStruct, err := parser.Parse(query)
		if err != nil {
			return err
		}
		list.Add(queryStruct)
	}

	for q := range list.Iter() {
		fmt.Fprintln(w, q)
	}

	return nil
}
