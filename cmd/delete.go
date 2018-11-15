package cmd

import (
	"fmt"
	"github.com/GaruGaru/done/tasks"
	"github.com/spf13/cobra"
	"time"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean tasks in the current day",
	Run: func(cmd *cobra.Command, args []string) {

		dbm, err := tasks.NewDBManager(DBPath())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dbm.Remove(time.Now())

	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
