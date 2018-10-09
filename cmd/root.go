package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sqd",
	Short: "sqd is suspicious query detection tool based on whitelist or blacklist",
	Long:  `sqd is suspicious query detection tool based on whitelist or blacklist`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
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
	cobra.OnInitialize()
}
