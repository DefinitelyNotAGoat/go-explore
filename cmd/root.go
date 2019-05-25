package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "Explorer",
		Short: "An efficient all in one block explorer",
	}

	rootCommand.AddCommand(
		newSync(),
	)

	return rootCommand
}

func Execute() error {
	return newRootCommand().Execute()
}
