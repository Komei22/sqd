package cmd

import (
	"fmt"
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

			list, err := lister.Create(os.Stdin)
			if err != nil {
				return err
			}

			for q := range list.Iter() {
				fmt.Fprintln(os.Stdout, q)
			}
			return nil
		},
	}

	return createCmd
}
