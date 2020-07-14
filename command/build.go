package command

import (
	"github.com/spf13/cobra"
	"gitlab.com/lcook/bugzport/internal/config"
	"gitlab.com/lcook/bugzport/internal/jail"
	"gitlab.com/lcook/bugzport/internal/port"
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
	buildCmd.Flags().BoolP("config", "c", false, "Configure port options before build (implies -o)")
}

var buildCmd = &cobra.Command{
	Use:   "build <category/name> [flags]",
	Short: "Build a port from a selected ports tree.",
	Long: `Trivially start the process of building a port in Poudriere
with a selected ports tree. You can use this to also generate
a summary report (diff, build log and port summary) once the
compliation process has finished. The report is a basis of what
to upload on the FreeBSD bugzilla and will speed up the committing
phase by following best practises.`,
	Example: `  # builds port with selected jail (-j), path to the ports directory (-d)
  # and uses 'devel' ports tree (-t)
  $ bp build devel/gh -j amd64-121-rel -d /path/to/ports -t devel

  # outputs build process (-o) and generates report (-r)
  $ bp build misc/broot -or

  # drops to an interactive shell (-i) inside the jail
  $ bp build lang/go -i`,
	Args: cobra.MinimumNArgs(1),
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

	jail, err := jail.FromName(jailName, treeName)
	if err != nil {
		return poudriere.Job{}, err
	}

	port, err := port.FromName(dirName + portName)
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
