package main

import (
	"os"

	"github.com/redhatrises/dockerfile-template/cmd/dft/common"
	_ "github.com/redhatrises/dockerfile-template/cmd/dft/generate"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd = parseCommands()

	Execute()
	os.Exit(0)
}

func parseCommands() *cobra.Command {
	for _, c := range common.Commands {
		parent := rootCmd
		if c.Parent != nil {
			parent = c.Parent
		}
		parent.AddCommand(c.Command)

		// - templates need to be set here, as PersistentPreRunE() is
		// not called when --help is used.
		// - rootCmd uses cobra default template not ours
		c.Command.SetHelpTemplate(helpTemplate)
		c.Command.SetUsageTemplate(usageTemplate)
		c.Command.DisableFlagsInUseLine = true
	}
	/*
		if err := terminal.SetConsole(); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
	*/
	return rootCmd
}
