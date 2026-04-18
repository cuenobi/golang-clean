package cmd

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "DDD + Clean Architecture reference service",
	}

	cmd.AddCommand(newAPICommand())
	cmd.AddCommand(newConsumerCommand())
	cmd.AddCommand(newMigrateCommand())

	return cmd
}
