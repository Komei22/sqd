package cmd

import (
	"os"

	"github.com/Komei22/sqd/lister"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func newCreateCmd() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "`sqd create` collect query and create whitelist",
		Long:  "`sqd create` collect query from stdin and create whitelist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if terminal.IsTerminal(0) {
				cmd.Help()
				return nil
			}
			lister.Create(os.Stdin, os.Stdout)

			return nil
		},
	}

	return createCmd
}
