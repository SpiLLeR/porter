package main

import (
	"fmt"

	"github.com/deislabs/porter/pkg/mixin/exec"

	"github.com/spf13/cobra"
)

func buildBuildCommand(m *exec.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(m.Out, "# exec mixin has no buildtime dependencies")
		},
	}
	return cmd
}
