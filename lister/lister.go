package lister

import (
	"bufio"
	"io"

	"github.com/Komei22/sqd/formatter"
	"github.com/Komei22/sql-mask"
	"github.com/deckarep/golang-set"
)

// Create whitelist
func Create(r io.Reader, m masker.Masker) (mapset.Set, error) {
	list := mapset.NewSet()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		query := scanner.Text()
		queryStruct, err := masker.Mask(m, query)
		queryStruct = formatter.Format(queryStruct)
		if err != nil {
			return nil, err
		}
		list.Add(queryStruct)
	}

	return list, nil
}
