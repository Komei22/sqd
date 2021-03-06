package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/eventor"
	"github.com/Komei22/sqd/matcher"
	"github.com/Komei22/sql-mask"
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
		database  string
	)

	rootCmd := &cobra.Command{
		Use:   "sqd",
		Short: "sqd is suspicious query detection tool based on whitelist or blacklist",
		Long:  `sqd is suspicious query detection tool based on whitelist or blacklist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var msk masker.Masker
			switch database {
			case "mysql":
				msk = &masker.MysqlMasker{}
			case "pg":
				msk = &masker.PgMasker{}
			default:
				return fmt.Errorf("Please set target database, `mysql` or `pg`")
			}

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

			d := detector.New(m, msk, mode)

			if querylog != "" {
				r, err := os.Open(querylog)
				if err != nil {
					return err
				}
				defer r.Close()
				detectQueries(r, d)
			} else if query != "" {
				suspiciousQuery, err := d.Detect(query)
				if err != nil {
					return fmt.Errorf("Can't detection suspicious query: %s", err)
				}
				fmt.Fprintln(os.Stdout, suspiciousQuery)
			} else {
				if terminal.IsTerminal(0) {
					cmd.Help()
					return nil
				}
				detectQueries(os.Stdin, d)
			}

			return nil
		},
	}
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "query string")
	rootCmd.Flags().StringVarP(&querylog, "file", "f", "", "query log file path")
	rootCmd.Flags().StringVarP(&whitelist, "Whitelist", "W", "", "whitelist file path")
	rootCmd.Flags().StringVarP(&blacklist, "Blacklist", "B", "", "blacklist file path")
	rootCmd.Flags().StringVarP(&database, "database", "d", "mysql", "target database")

	rootCmd.AddCommand(newCreateCmd())

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

func detectQueries(r io.Reader, d *detector.Detector) {
	suspiciousQueryChan := make(chan string)
	errChan := make(chan error)
	defer close(errChan)
	defer close(suspiciousQueryChan)

	go d.DetectFrom(r, suspiciousQueryChan, errChan)
	eventor.Print(os.Stdout, suspiciousQueryChan, errChan)
}
