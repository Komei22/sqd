package lister

import (
	"bufio"
	"io"

	"github.com/Komei22/sqd/formatter"
	"github.com/Komei22/sql-mask"
	"github.com/deckarep/golang-set"
)

// Create whitelist
func Create(r io.Reader) (mapset.Set, error) {
	list := mapset.NewSet()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		query := formatter.Format(scanner.Text())
		queryStruct, err := parser.Parse(query)
		if err != nil {
			return nil, err
		}
		list.Add(queryStruct)
	}

	return list, nil
}
