package cmd

import (
	"domain-stress/app"
	"errors"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "domain-stress",
	Short: "stress to the domian",
	Long:  `the type to stress to the domian .`,
	Run: func(cmd *cobra.Command, args []string) {
		app.Error(cmd, args, errors.New("error command"))
	},
}

func Execute() {
	runtime.GOMAXPROCS(1)
	rootCmd.Execute()
}
