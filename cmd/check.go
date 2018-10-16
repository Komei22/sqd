package cmd

import (
	"fmt"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/matcher"
	"github.com/spf13/cobra"
)

var (
	query            string
	querylogFilepath string
	listFilepath     string
	isWhitelistMode  bool
	isBlacklistMode  bool
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "`sqd check` investigate query base on white/black list",
	Long:  "`sqd check` investigate query base on white/black list",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := matcher.New(listFilepath)
		if err != nil {
			fmt.Printf("Can't read list file. (%s)", err)
			os.Exit(1)
		}

		var mode detector.Mode
		if (!isWhitelistMode && !isBlacklistMode) || (isWhitelistMode && isBlacklistMode) {
			fmt.Print("Please set detection mode, Whitelist(-W) or Blacklist(-B)")
			os.Exit(1)
		} else if !isWhitelistMode && isBlacklistMode {
			mode = detector.Blacklist
		} else if isWhitelistMode && !isBlacklistMode {
			mode = detector.Whitelist
		}

		d := detector.New(m, mode)
		if err != nil {
			fmt.Printf("Can't read input query. (%s)", err)
			os.Exit(1)
		}

		var suspiciousQueries []string
		if querylogFilepath != "" {
			suspiciousQueries, err = d.DetectFrom(querylogFilepath)
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

func init() {
	checkCmd.Flags().StringVarP(&query, "query", "q", "", "query string")
	checkCmd.Flags().StringVarP(&querylogFilepath, "file", "f", "", "query log file path")
	checkCmd.Flags().BoolVarP(&isWhitelistMode, "Whitelist", "W", false, "run whitelist mode")
	checkCmd.Flags().BoolVarP(&isBlacklistMode, "Blacklist", "B", false, "run blacklist mode")
	checkCmd.Flags().StringVarP(&listFilepath, "list", "l", "", "file path of blacklist or whitelist")

	rootCmd.AddCommand(checkCmd)
}
