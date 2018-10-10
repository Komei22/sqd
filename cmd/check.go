package cmd

import (
	"fmt"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/matcher"
	"github.com/spf13/cobra"
)

var (
	querylogFilepath  string
	whitelistFilepath string
	blacklistFilepath string
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "`sqd check` investigate query log base on white/black list",
	Long:  "`sqd check` investigate query log base on white/black list",
	Run: func(cmd *cobra.Command, args []string) {

		m := &matcher.Matcher{}
		var err error
		var detectionMode detector.DetectionMode
		if whitelistFilepath != "" {
			m, err = matcher.New(whitelistFilepath)
			if err != nil {
				// logger.Warn("Can't read list file", zap.Error(err))
				fmt.Print(err)
				os.Exit(1)
			}
			detectionMode = detector.Whitelist
		} else {
			m, err = matcher.New(blacklistFilepath)
			if err != nil {
				// logger.Warn("Can't read list file", zap.Error(err))
				fmt.Print(err)
				os.Exit(1)
			}
			detectionMode = detector.Blacklist
		}

		d, err := detector.New(querylogFilepath, detectionMode)
		if err != nil {
			// logger.Warn("Can't read query log file", zap.Error(err))
			fmt.Print(err)
			os.Exit(1)
		}

		d.DumpSuspiciousQuerys(m)
	},
}

func init() {
	checkCmd.Flags().StringVarP(&querylogFilepath, "queryfile", "q", "", "query log file path")
	checkCmd.Flags().StringVarP(&whitelistFilepath, "whitelist", "w", "", "whitelist file path")
	checkCmd.Flags().StringVarP(&blacklistFilepath, "blacklist", "b", "", "blacklist file path")

	rootCmd.AddCommand(checkCmd)
}
