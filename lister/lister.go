package lister

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Komei22/sql-mask"
	"github.com/deckarep/golang-set"
)

// Create whitelist
func Create(r io.Reader, output string) error {
	list := mapset.NewSet()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		query := scanner.Text()
		fomatedQuery := trimControlChara(query)
		fmt.Println(fomatedQuery)
		queryStruct, err := parser.Parse(query)
		if err != nil {
			return err
		}
		list.Add(queryStruct)
	}

	return save(output, list)
}

func save(filepath string, list mapset.Set) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	it := list.Iterator()

	for q := range it.C {
		file.Write(([]byte)(q.(string) + "\n"))
	}
	return nil
}

func trimControlChara(query string) string {
	query = query[1:strings.LastIndex(query, "\"")]
	query = strings.Replace(query, "\\n", " ", -1)
	query = strings.Replace(query, "\\", "", -1)
	return query
}
