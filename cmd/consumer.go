package cmd

import (
	"context"
	"fmt"

	"github.com/cuenobi/golang-clean/internal/bootstrap"
	"github.com/spf13/cobra"
)

func newConsumerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "Run Kafka consumers",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := bootstrap.NewApp()
			if err != nil {
				return err
			}
			defer app.Close(context.Background())

			fmt.Println("Kafka consumer started")
			return app.RunConsumer(context.Background())
		},
	}
}
