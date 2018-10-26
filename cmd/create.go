package cmd

import (
	"bufio"
	"os"

	"github.com/Komei22/sql-mask"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func newCreateCmd() *cobra.Command {
	var filepath string

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "`sqd create` collect query and create whitelist",
		Long:  "`sqd create` collect query from stdin and create whitelist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if terminal.IsTerminal(0) {
				cmd.Help()
				return nil
			}

			file, err := os.Create(filepath)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				if err := scanner.Err(); err != nil {
					return err
				}
				query := scanner.Text()
				queryStruct, err := parser.Parse(query)
				if err != nil {
					return err
				}
				file.Write(([]byte)(queryStruct + "\n"))
			}

			return nil
		},
	}

	createCmd.Flags().StringVarP(&filepath, "output", "o", "whitelist", "output path of whitelist file(default: ./whitelist)")

	return createCmd
}
