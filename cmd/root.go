package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   APP,
		Short: "Manage stacks",
		Long:  `Command line tool to configure and deploy aws cloudformation stacks`,
		Run: func(cmd *cobra.Command, args []string) {
			if run.version {
				versionCmd.Run(cmd, args)
				return
			}

			cmd.Help()
		},
	}
)

func Execute(version string) {
	VERSION = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&run.version, "version", "v", false, "print current/running version")
	rootCmd.AddCommand(versionCmd)
}
