package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"matcha/internal/config"
	"matcha/internal/driver/api"
)

func main() {
	cmd := NewCmdRoot()
	cmd.AddCommand(NewCmdServer())
	if err := cmd.Execute(); err != nil {
		log.Error().Err(err).Msg("failed to execute command")
		os.Exit(1)
	}
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "matcha",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if config.Get().IsDevelopment() {
				log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			}
			return nil
		},
	}
	return cmd
}

func NewCmdServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return api.Run(cmd.Context())
		},
	}
	return cmd
}
