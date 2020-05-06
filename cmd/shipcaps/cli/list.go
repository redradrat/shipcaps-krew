package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	listCmd.AddCommand(listAppCmd)
	listCmd.AddCommand(listCapCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shipcaps resources",
	Long:  `This command is used to search for and list shipcaps resources.`,
}

var listAppCmd = &cobra.Command{
	Use:     "app",
	Aliases: []string{"apps"},
	Short:   "List deployed apps",
	Long:    `List deployed apps in the cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing apps...")
	},
}

var listCapCmd = &cobra.Command{
	Use:     "cap",
	Aliases: []string{"caps"},
	Short:   "List deployed caps",
	Long:    `List deployed caps in the cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing caps...")
	},
}
