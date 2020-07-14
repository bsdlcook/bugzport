package command

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "bp <command> [flags]",
	Long: `bugzport, the friendly wrapper for Poudriere that allows
you to build ports and generate a summary report for
easier submissions on bugzilla. This tool is handy and
aimed at FreeBSD port maintainers.`,
	SilenceErrors: true,
	SilenceUsage:  true,
}
