package cmd

import (
	"flemming-pr/gray/server"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gray",
	Short: "Start the gray server",
	Long:  `Start the gray server to render incoming messages`,
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
