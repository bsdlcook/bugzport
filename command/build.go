package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/lcook/bugzport/internal/config"
	"gitlab.com/lcook/bugzport/internal/poudriere"
)

func init() {
	RootCmd.AddCommand(buildCmd)

	cfg, _ := config.Get()
	buildCmd.Flags().StringP("dir", "d", cfg.Dir, "Target ports directory")
	buildCmd.Flags().StringP("jail", "j", cfg.Jail, "Target jail")
	buildCmd.Flags().StringP("tree", "t", cfg.Tree, "Target ports tree")
	buildCmd.Flags().BoolP("report", "r", false, "Generate a report once finished building")
	buildCmd.Flags().BoolP("output", "o", false, "Show running output of Poudriere build process")
}

var buildCmd = &cobra.Command{
	Use:   "build <category/name> [flags]",
	Short: "Queue a port for a build in Poudriere.",
	Long: `Queue port for a build in Poudriere, generate a report and diff of the changes made.

Examples:
  $ bp build audio/spotify-tui -j builder-amd64-121-rel -d /path/to/ports -t default
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			return fmt.Errorf("Port is required as argument: <category/name>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		portName := args[0]

		dirName, _ := cmd.Flags().GetString("dir")
		jailName, _ := cmd.Flags().GetString("jail")

		tree, _ := cmd.Flags().GetString("tree")
		output, _ := cmd.Flags().GetBool("output")
		report, _ := cmd.Flags().GetBool("report")

		jail, err := poudriere.JailFromName(jailName, tree)
		if err != nil {
			return err
		}

		port, err := poudriere.PortFromName(dirName + portName)
		if err != nil {
			return err
		}

		options := &poudriere.OptionsT{
			Output: output,
			Report: report,
		}

		job := poudriere.Job{
			Jail:    jail,
			Port:    port,
			Tree:    tree,
			WorkDir: dirName,
			Options: options,
		}

		job.Run()
		return nil
	},
}
