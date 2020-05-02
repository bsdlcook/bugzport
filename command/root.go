package command

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "bugzport <command> [flags]",
	Long:          "Wrapper around Poudriere which allows you to build ports and generate a summary report for submissions on bugzilla. This tool is handy and aimed at FreeBSD port maintainers.",
	SilenceErrors: true,
	SilenceUsage:  true,
}
