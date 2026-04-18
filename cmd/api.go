package cmd

import (
	"context"
	"fmt"

	"github.com/cuenobi/golang-clean/internal/bootstrap"
	"github.com/spf13/cobra"
)

func newAPICommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Run HTTP API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := bootstrap.NewApp()
			if err != nil {
				return err
			}
			defer app.Close(context.Background())

			fmt.Printf("HTTP API listening on %s\n", app.Config.HTTPAddress)
			return app.RunHTTP()
		},
	}
}
