package cmd

import (
	"fmt"
	"strconv"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

func newMigrateCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "migrate", Short: "Run database migrations"}

	cmd.AddCommand(&cobra.Command{
		Use:   "up",
		Short: "Apply all up migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := newMigrator()
			if err != nil {
				return err
			}
			defer func() { _, _ = m.Close() }()
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				return err
			}
			return nil
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "down [steps]",
		Short: "Rollback migration steps",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := newMigrator()
			if err != nil {
				return err
			}
			defer func() { _, _ = m.Close() }()

			if len(args) == 0 {
				return m.Down()
			}
			steps, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid steps: %w", err)
			}
			return m.Steps(-steps)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print migration version",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := newMigrator()
			if err != nil {
				return err
			}
			defer func() { _, _ = m.Close() }()

			version, dirty, err := m.Version()
			if err == migrate.ErrNilVersion {
				fmt.Println("no migration has been applied")
				return nil
			}
			if err != nil {
				return err
			}
			fmt.Printf("version=%d dirty=%t\n", version, dirty)
			return nil
		},
	})

	return cmd
}

func newMigrator() (*migrate.Migrate, error) {
	cfg := config.Load()
	sourceURL := "file://migrations"
	m, err := migrate.New(sourceURL, cfg.PostgresMigrationURL())
	if err != nil {
		return nil, err
	}
	return m, nil
}
