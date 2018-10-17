package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRootCmdNormal(t *testing.T) {
	cases := []struct {
		command string
		want    string
		cmdArgs []string
	}{
		{command: "sqd", cmdArgs: []string{"-q", "DROP TABLE articles", "-B", "../testdata/blacklist"}, want: "Suspicious queries:\nDROP TABLE articles\n"},
		{command: "sqd", cmdArgs: []string{"-q", "DROP TABLE articles", "-W", "../testdata/whitelist"}, want: "Suspicious queries:\nDROP TABLE articles\n"},
		{command: "sqd", cmdArgs: []string{"-f", "../testdata/query.log", "-B", "../testdata/blacklist"}, want: "Suspicious queries:\nSELECT * FROM articles\nDROP TABLE articles\n"},
		{command: "sqd", cmdArgs: []string{"-f", "../testdata/query.log", "-W", "../testdata/whitelist"}, want: "Suspicious queries:\nSELECT * FROM articles\nDROP TABLE articles\n"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := newRootCmd()
		cmd.SetOutput(buf)
		cmd.SetArgs(c.cmdArgs[:])
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			fmt.Printf("cmdArgs %s\n", c.cmdArgs[:])
			t.Errorf("unexpected response: want:%s, get:%s", c.want, get)
		}
	}
}
