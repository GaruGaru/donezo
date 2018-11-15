package cmd

import (
	"github.com/spf13/cobra"
)


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "print tasks list",
}

func init()  {
	rootCmd.AddCommand(listCmd)
}