package command

import (
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
	Use: "build <category/name> [flags]",
	//Short: "",
	//Long:  "",
	Args: cobra.MinimumNArgs(1),
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

		jail := poudriere.JailFromName(jailName)
		port := poudriere.PortFromName(dirName + portName)

		job := poudriere.Job{
			Jail: jail,
			Port: port,
			Tree: tree,
		}

		job.Run()
		return nil
	},
}
