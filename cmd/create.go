package cmd

import (
	"os"

	"github.com/Komei22/sqd/lister"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func newCreateCmd() *cobra.Command {
	var output string

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "`sqd create` collect query and create whitelist",
		Long:  "`sqd create` collect query from stdin and create whitelist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if terminal.IsTerminal(0) {
				cmd.Help()
				return nil
			}
			lister.Create(os.Stdin, output)

			return nil
		},
	}

	createCmd.Flags().StringVarP(&output, "output", "o", "whitelist", "output path of whitelist file(default: ./whitelist)")

	return createCmd
}
