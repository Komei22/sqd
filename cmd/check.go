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

var (
	querylogFilepath string
	listFilepath     string
	isWhitelistMode  bool
	isBlacklistMode  bool
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
		if (!isWhitelistMode && !isBlacklistMode) || (isWhitelistMode && isBlacklistMode) {
			fmt.Print("Please set detection mode, Whitelist(-W) or Blacklist(-B)")
			os.Exit(1)
		} else if !isWhitelistMode && isBlacklistMode {
			mode = detector.Blacklist
		} else if isWhitelistMode && !isBlacklistMode {
			mode = detector.Whitelist
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
	checkCmd.Flags().BoolVarP(&isWhitelistMode, "Whitelist", "W", false, "run whitelist mode")
	checkCmd.Flags().BoolVarP(&isBlacklistMode, "Blacklist", "B", false, "run blacklist mode")
	checkCmd.Flags().StringVarP(&listFilepath, "list", "l", "", "file path of blacklist or whitelist")

	rootCmd.AddCommand(checkCmd)

}
