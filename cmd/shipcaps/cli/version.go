package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/redradrat/shipcaps-kubectl/pkg/version"
)

func init() {
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the shipcaps-kubectl version",
	Long:  `This command is used to get version information about the kubectl shipcaps shipcaps.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Git Version: %s", version.GitTag()))
		fmt.Println(fmt.Sprintf("Git Commit: %s", version.GitCommit()))
	},
}
