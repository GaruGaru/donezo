package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean tasks in the current day",
	Run: func(cmd *cobra.Command, args []string) {
		repository := Repository()
		repository.Remove(time.Now())
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
