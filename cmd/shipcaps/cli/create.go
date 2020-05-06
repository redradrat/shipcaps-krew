package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(createAppCmd)
	createCmd.AddCommand(createCapCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create shipcaps resource",
	Long:  `This command is used to assist in creating a valid shipcaps resources with a guide.`,
}

var createAppCmd = &cobra.Command{
	Use:     "app",
	Aliases: []string{"apps"},
	Short:   "Create an app based on a cap resource",
	Long:    `This command is used to assist in creating a valid shipcaps app from a deployed cap.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating app...")
	},
}

var createCapCmd = &cobra.Command{
	Use:     "cap",
	Aliases: []string{"caps"},
	Short:   "Create a cap resource",
	Long:    `This command is used to assist in creating a valid shipcaps cap.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating cap...")
	},
}
