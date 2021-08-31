package main

import (
	"fmt"

	"github.com/redhatrises/dockerfile-template/cmd/dft/common"
	"github.com/redhatrises/dockerfile-template/version"
	"github.com/spf13/cobra"
)

var (
	// Command: dft _version_
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version for dft",
		Long:  "version for dft",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("dft version %s\n", version.Version.String())
			return nil
		},
	}
)

func init() {
	common.Commands = common.AllCommands(common.CliCommand{
		Command: versionCmd,
	})
}
