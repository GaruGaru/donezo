package cmd

import (
	"fmt"
	"github.com/GaruGaru/done/tasks"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var (
	taskEffort     int
	taskEffortType string
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "add a task for the current day",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Print("no task provided")
			return
		}

		taskDescription := strings.Join(args, " ")

		dbm, err := tasks.NewDBManager(DBPath())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dbm.Add(time.Now(), tasks.Task{
			Description: taskDescription,
			Effort:      &taskEffort,
			EffortType:  &taskEffortType,
		})

	},
}

func init() {
	addCmd.Flags().StringVarP(&taskEffortType, "effort-type", "t", "", "effor type for the task (hours, points...)")
	addCmd.Flags().IntVarP(&taskEffort, "effort", "e", 0, "effor value for the task (hours, points...)")
	rootCmd.AddCommand(addCmd)
}
