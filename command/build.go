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
		dirName, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}

		jailName, err := cmd.Flags().GetString("jail")
		if err != nil {
			return err
		}

		tree, err := cmd.Flags().GetString("tree")
		if err != nil {
			return err
		}

		jail, err := poudriere.JailFromName(jailName, tree)
		if err != nil {
			return err
		}

		port, err := poudriere.PortFromName(dirName + portName)
		if err != nil {
			return err
		}

		job := poudriere.Job{
			Jail:    jail,
			Port:    port,
			Tree:    tree,
			WorkDir: dirName,
		}

		job.Run()
		return nil
	},
}
