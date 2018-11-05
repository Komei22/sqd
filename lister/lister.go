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
		in := scanner.Text()
		query, err := formatter.ExtractQueryFrom(in)
		if err != nil {
			return nil, err
		}
		queryStruct, err := parser.Parse(query)
		queryStruct = formatter.Format(queryStruct)
		if err != nil {
			return nil, err
		}
		list.Add(queryStruct)
	}

	return list, nil
}
