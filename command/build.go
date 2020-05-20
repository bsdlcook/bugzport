package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/lcook/bugzport/internal/config"
	"gitlab.com/lcook/bugzport/internal/poudriere"
)

func init() {
	RootCmd.AddCommand(buildCmd)

	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	buildCmd.Flags().StringP("dir", "d", cfg.Dir, "Target ports directory")
	buildCmd.Flags().StringP("jail", "j", cfg.Jail, "Target jail")
	buildCmd.Flags().StringP("tree", "t", cfg.Tree, "Target ports tree")
	buildCmd.Flags().BoolP("report", "r", false, "Generate a report once finished building")
	buildCmd.Flags().BoolP("output", "o", false, "Show running output of Poudriere build process")
	buildCmd.Flags().BoolP("interactive", "i", false, "Start an interactive shell inside the builder once built (implies -o)")
	buildCmd.Flags().BoolP("config", "c", false, "Configure port options before build. (implies -o)")
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
			return fmt.Errorf("port is required as argument: <category/name>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		buildJob, err := poudriereJob(cmd, args)
		if err != nil {
			return err
		}

		buildJob.Run()
		return nil
	},
}

func poudriereJob(cmd *cobra.Command, args []string) (poudriere.Job, error) {
	portName := args[0]
	dirName := getString(cmd, "dir")
	jailName := getString(cmd, "jail")
	treeName := getString(cmd, "tree")

	jail, err := poudriere.JailFromName(jailName, treeName)
	if err != nil {
		return poudriere.Job{}, err
	}

	port, err := poudriere.PortFromName(dirName + portName)
	if err != nil {
		return poudriere.Job{}, err
	}

	options := getOptions(cmd)

	return poudriere.Job{
		Jail:    jail,
		Port:    port,
		Tree:    treeName,
		WorkDir: dirName,
		Options: options,
	}, nil
}

func getBool(cmd *cobra.Command, value string) bool {
	val, _ := cmd.Flags().GetBool(value)
	return val
}

func getString(cmd *cobra.Command, value string) string {
	val, _ := cmd.Flags().GetString(value)
	return val
}

func getOptions(cmd *cobra.Command) *poudriere.Options {
	return &poudriere.Options{
		Output:      getBool(cmd, "output"),
		Report:      getBool(cmd, "report"),
		Interactive: getBool(cmd, "interactive"),
		Config:      getBool(cmd, "config"),
	}
}
