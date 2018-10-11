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
	detectorMode     string
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

		var mode detector.Mode
		switch detectorMode {
		case "whitelist":
			mode = detector.Whitelist
		case "blacklist":
			mode = detector.Blacklist
		default:
			fmt.Printf("Unknown detection mode.(%s)", detectorMode)
			os.Exit(1)
		}
		d, err := detector.New(querylogFilepath, mode)
		if err != nil {
			fmt.Printf("Can't read query log file. (%s)", err)
			os.Exit(1)
		}

		suspiciousQuerys, err := d.Detect(m)
		if err != nil {
			fmt.Printf("Can't detection suspicious query. (%s)", err)
			os.Exit(1)
		}

		fmt.Print("Suspicious querys\n")
		for _, query := range suspiciousQuerys {
			fmt.Printf("%s\n", query)
		}
	},
}

func init() {
	checkCmd.Flags().StringVarP(&querylogFilepath, "queryfile", "q", "", "query log file path")
	checkCmd.Flags().StringVarP(&listFilepath, "list", "l", "", "file path of blacklist or whitelist")
	checkCmd.Flags().StringVarP(&detectorMode, "mode", "m", "whitelist", "detection mode (whitelist or blacklist)")

	rootCmd.AddCommand(checkCmd)
}
