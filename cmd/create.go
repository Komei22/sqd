package cmd

import (
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "`sqd create` collect query and create whitelist",
		Long:  "`sqd create` collect query and create whitelist",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return createCmd
}
