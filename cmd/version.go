package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show domain-stress version",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version %s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
