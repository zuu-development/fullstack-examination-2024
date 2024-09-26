package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zuu-development/fullstack-examination-2024/internal/common"
)

func init() {
	rootCmd.AddCommand(NewVersionCmd())
}

// NewVersionCmd returns a new `version` command to be used as a sub-command to root
func NewVersionCmd() *cobra.Command {
	var output string

	versionCmd := cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Example: `  # Print the full version of client and server to stdout
  todo-cli version
`,
		Run: func(cmd *cobra.Command, _ []string) {
			cv := common.GetVersion()
			switch output {
			case "wide", "short", "":
				fmt.Fprint(cmd.OutOrStdout(), printClientVersion(&cv, (output == "short")))
			default:
				log.Fatalf("unknown output format: %s", output)
			}
		},
	}
	versionCmd.Flags().StringVarP(&output, "output", "o", "wide", "Output format. One of: wide|short")
	return &versionCmd
}

func printClientVersion(version *common.Version, short bool) string {
	output := fmt.Sprintf("%s: %s\n", "todo-cli", version)
	if short {
		return output
	}
	output += fmt.Sprintf("  BuildDate: %s\n", version.BuildDate)
	output += fmt.Sprintf("  GitCommit: %s\n", version.GitCommit)
	output += fmt.Sprintf("  GoVersion: %s\n", version.GoVersion)
	output += fmt.Sprintf("  Compiler: %s\n", version.Compiler)
	output += fmt.Sprintf("  Platform: %s\n", version.Platform)
	return output
}
