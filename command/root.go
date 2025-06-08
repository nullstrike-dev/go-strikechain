package command

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "strikechain",
	Short: "StrikeChain CLI",
	Long:  `Command line interface for StrikeChain blockchain`,
}

func Execute() {
	rootCmd.Execute()
}
