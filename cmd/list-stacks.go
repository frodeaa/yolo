package cmd

import (
	"github.com/frodeaa/yolo/cfn"
	"github.com/spf13/cobra"
)

var listStacks = &cobra.Command{
	Use:   "list-stacks",
	Short: "list all stacks",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cfn.NewContext()

		res, err := ctx.ListStacks()

		if err != nil {
			// TODO handle err
			panic(err)
			return
		}

		printer := cfn.NewStackPrinter(run.monocromeOutput, cmd.OutOrStderr())
		printer.PrintStacks(res)
	},
}

func init() {
	listStacks.Flags().BoolVarP(&run.monocromeOutput, "monochrome-output", "m", false, "disable color output")
}
