package main

import (
	"fmt"
	"os"

	"github.com/redhatrises/dockerfile-template/cmd/dft/common"
	"github.com/redhatrises/dockerfile-template/version"
	"github.com/spf13/cobra"
)

// HelpTemplate is the help template
// This uses the short and long options.
// command should not use this.
const helpTemplate = `{{.Short}}
Description:
  {{.Long}}
{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

// UsageTemplate is the usage template
// This blocks the displaying of the global options. The main
// command should not use this.
const usageTemplate = `Usage:{{if (and .Runnable (not .HasAvailableSubCommands))}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.UseLine}} [command]{{end}}{{if gt (len .Aliases) 0}}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}
Examples:
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
Options:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
{{end}}
`

var (
	rootCmd = &cobra.Command{
		Use:                   "dft",
		Short:                 "dft - Create Dockerfiles using Go templating",
		Long:                  `A tool to create multiple Dockerfiles using Go-style templating`,
		SilenceUsage:          false,
		SilenceErrors:         true,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		Version:               version.Version.String(),
		RunE:                  common.SubCommandExists,
		PersistentPreRunE:     persistentPreRunE,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// For anything related to Cobra initializing
}

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
