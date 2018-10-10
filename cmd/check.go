package cmd

import (
	"fmt"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/matcher"
	"github.com/spf13/cobra"
)

var (
	querylogFilepath string
	listFilepath     string
	detectionMode    string
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "`sqd check` investigate query log base on white/black list",
	Long:  "`sqd check` investigate query log base on white/black list",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := matcher.New(listFilepath)
		if err != nil {
			fmt.Printf("Can't read list file. (%s)", err)
			os.Exit(1)
		}

		d, err := detector.New(querylogFilepath, detectionMode)
		if err != nil {
			fmt.Printf("Can't read query log file. (%s)", err)
			os.Exit(1)
		}

		suspiciousQuerys, err := d.Detect(m)
		if err != nil {
			fmt.Printf("Can't detection suspicious query. (%s)", err)
			os.Exit(1)
		}
		detector.Dump(suspiciousQuerys)
	},
}

func init() {
	checkCmd.Flags().StringVarP(&querylogFilepath, "queryfile", "q", "", "query log file path")
	checkCmd.Flags().StringVarP(&listFilepath, "list", "l", "", "file path of blacklist or whitelist")
	checkCmd.Flags().StringVarP(&detectionMode, "mode", "m", "whitelist", "detection mode (whitelist or blacklist)")

	rootCmd.AddCommand(checkCmd)
}
