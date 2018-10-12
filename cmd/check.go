package cmd

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/matcher"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "`sqd check` investigate query base on white/black list",
	Long:  "`sqd check` investigate query base on white/black list",
	Args: func(cmd *cobra.Command, args []string) error {
		if terminal.IsTerminal(0) {
			if len(args) < 1 {
				return errors.New("requires [query ...]")
			}
		}
		return nil
	},
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
		d, err := detector.New(args, mode)
		if err != nil {
			fmt.Printf("Can't read input query. (%s)", err)
			os.Exit(1)
		}

		suspiciousQueries, err := d.Detect(m)
		if err != nil {
			fmt.Printf("Can't detection suspicious query. (%s)", err)
			os.Exit(1)
		}

		fmt.Print("Suspicious queries:\n")
		for _, query := range suspiciousQueries {
			fmt.Printf("%s\n", query)
		}
	},
}

func init() {
	checkCmd.Flags().StringVarP(&listFilepath, "list", "l", "", "file path of blacklist or whitelist")
	checkCmd.Flags().StringVarP(&detectorMode, "mode", "m", "whitelist", "detection mode (whitelist or blacklist)")

	rootCmd.AddCommand(checkCmd)

}
