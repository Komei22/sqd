package cmd

import (
	"fmt"
	"os"

	"github.com/Komei22/sqd/lister"
	"github.com/Komei22/sql-mask"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func newCreateCmd() *cobra.Command {
	var database string

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "`sqd create` collect query and create whitelist",
		Long:  "`sqd create` collect query from stdin and create whitelist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if terminal.IsTerminal(0) {
				cmd.Help()
				return nil
			}

			var m masker.Masker
			switch database {
			case "mysql":
				m = &masker.MysqlMasker{}
			case "pg":
				m = &masker.PgMasker{}
			default:
				return fmt.Errorf("Please set target database, `mysql` or `pg`")
			}

			list, err := lister.Create(os.Stdin, m)
			if err != nil {
				return err
			}

			for q := range list.Iter() {
				fmt.Fprintln(os.Stdout, q)
			}
			return nil
		},
	}

	createCmd.Flags().StringVarP(&database, "database", "d", "mysql", "target database")

	return createCmd
}
