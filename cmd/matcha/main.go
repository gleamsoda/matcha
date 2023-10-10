package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"matcha/internal/config"
)

func main() {
	cmd := NewCmdRoot()
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
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info().Msg("Hello, world!")
			return errors.New("not implemented")
		},
	}
	return cmd
}
