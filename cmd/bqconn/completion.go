// Code generated. DO NOT EDIT.

package main

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Emits bash a completion for bqconn",
	Long: `Enable bash completion like so:
		Linux:
			source <(bqconn completion)
		Mac:
			brew install bash-completion
			bqconn completion > $(brew --prefix)/etc/bash_completion.d/bqconn`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}
