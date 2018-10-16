package cmd

import (
	"fmt"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/matcher"
	"github.com/spf13/cobra"
)

var (
	query     string
	querylog  string
	blacklist string
	whitelist string
)

var rootCmd = &cobra.Command{
	Use:   "sqd",
	Short: "sqd is suspicious query detection tool based on whitelist or blacklist",
	Long:  `sqd is suspicious query detection tool based on whitelist or blacklist`,
	Run: func(cmd *cobra.Command, args []string) {
		var mode detector.Mode
		var m *matcher.Matcher
		var err error
		if (whitelist == "" && blacklist == "") || (whitelist != "" && blacklist != "") {
			fmt.Print("Please set list file path, Whitelist(-W) or Blacklist(-B)")
			os.Exit(1)
		} else if whitelist == "" && blacklist != "" {
			mode = detector.Blacklist
			m, err = matcher.New(blacklist)
			if err != nil {
				fmt.Printf("Can't read list file. (%s)", err)
				os.Exit(1)
			}
		} else if whitelist != "" && blacklist == "" {
			mode = detector.Whitelist
			m, err = matcher.New(whitelist)
			if err != nil {
				fmt.Printf("Can't read list file. (%s)", err)
				os.Exit(1)
			}
		}

		d := detector.New(m, mode)

		var suspiciousQueries []string
		if querylog != "" {
			suspiciousQueries, err = d.DetectFrom(querylog)
			if err != nil {
				fmt.Printf("Can't detection suspicious query. (%s)", err)
			}
		} else {
			suspiciousQuery, err := d.Detect(query)
			if err != nil {
				fmt.Printf("Can't detection suspicious query. (%s)", err)
				os.Exit(1)
			}
			suspiciousQueries = append(suspiciousQueries, suspiciousQuery)
		}

		fmt.Print("Suspicious queries:\n")
		for _, sq := range suspiciousQueries {
			fmt.Printf("%s\n", sq)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "query string")
	rootCmd.Flags().StringVarP(&querylog, "file", "f", "", "query log file path")
	rootCmd.Flags().StringVarP(&whitelist, "Whitelist", "W", "", "whitelist file path")
	rootCmd.Flags().StringVarP(&blacklist, "Blacklist", "B", "", "blacklist file path")

	cobra.OnInitialize()
}
