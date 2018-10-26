package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/eventor"
	"github.com/Komei22/sqd/matcher"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// NewRootCmd return rootCmd
func newRootCmd() *cobra.Command {
	var (
		query     string
		querylog  string
		blacklist string
		whitelist string
	)

	rootCmd := &cobra.Command{
		Use:   "sqd",
		Short: "sqd is suspicious query detection tool based on whitelist or blacklist",
		Long:  `sqd is suspicious query detection tool based on whitelist or blacklist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var mode detector.Mode
			var m *matcher.Matcher
			var err error
			if (whitelist == "" && blacklist == "") || (whitelist != "" && blacklist != "") {
				return errors.New("Please set list file path, Whitelist(-W) or Blacklist(-B)")
			} else if whitelist == "" && blacklist != "" {
				mode = detector.Blacklist
				m, err = matcher.New(blacklist)
				if err != nil {
					return fmt.Errorf("Can't read blacklist file: %s", err)
				}
			} else if whitelist != "" && blacklist == "" {
				mode = detector.Whitelist
				m, err = matcher.New(whitelist)
				if err != nil {
					return fmt.Errorf("Can't read whitelist file: %s", err)
				}
			}

			d := detector.New(m, mode)

			if querylog != "" {
				r, err := os.Open(querylog)
				if err != nil {
					return err
				}
				defer r.Close()
				err = detectQueries(r, d)
				if err != nil {
					return err
				}
			} else if query != "" {
				suspiciousQuery, err := d.Detect(query)
				if err != nil {
					return fmt.Errorf("Can't detection suspicious query: %s", err)
				}
				cmd.Println(suspiciousQuery)
			} else {
				if terminal.IsTerminal(0) {
					cmd.Help()
					return nil
				}
				err = detectQueries(os.Stdin, d)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "query string")
	rootCmd.Flags().StringVarP(&querylog, "file", "f", "", "query log file path")
	rootCmd.Flags().StringVarP(&whitelist, "Whitelist", "W", "", "whitelist file path")
	rootCmd.Flags().StringVarP(&blacklist, "Blacklist", "B", "", "blacklist file path")

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := newRootCmd()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func detectQueries(r io.Reader, d *detector.Detector) error {
	suspiciousQueryChan := make(chan string)
	errChan := make(chan error)
	go d.DetectFrom(r, suspiciousQueryChan, errChan)
	err := eventor.Print(os.Stdout, suspiciousQueryChan, errChan)
	if err != nil {
		return fmt.Errorf("Can't detection suspicious query: %s", err)
	}
	return nil
}
