package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show " + APP + " version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(APP + " " + VERSION)
	},
}
