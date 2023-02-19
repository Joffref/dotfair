package cmd

import (
	"github.com/dotfair-opensource/dotfair/pkg/dotfair"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	workspace string // workspace is the folder to scan containing terraform code
	verbose   bool   // verbose is the flag to enable verbose output
	format    string // format is the format of the output
	runCmd    = &cobra.Command{
		Use:   "run",
		Short: "Run a scan of your terraform code",
		Long:  `Run a scan of your terraform code`,
		RunE: func(cmd *cobra.Command, args []string) error {
			runner, err := dotfair.NewRunner(&dotfair.Config{
				Folder:       workspace,
				Verbose:      verbose,
				OutputFormat: format,
			})
			if err != nil {
				return err
			}
			return runner.Run(cmd.Context())
		},
	}
)

func init() {
	runCmd.Flags().StringVarP(&workspace, "workspace", "w", "./terraform", "Specify the folder to scan containing terraform code")
	runCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	runCmd.Flags().StringVarP(&format, "format", "f", "human-readable", "Specify the format of the output, either human-readable or json")
	viper.BindPFlag("workspace", runCmd.Flags().Lookup("workspace"))
	viper.BindPFlag("verbose", runCmd.Flags().Lookup("verbose"))
	viper.BindPFlag("format", runCmd.Flags().Lookup("format"))
	rootCmd.AddCommand(runCmd)
}
