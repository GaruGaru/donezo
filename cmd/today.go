package cmd

import (
	"fmt"
	"github.com/GaruGaru/done/tasks"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "list tasks for the current day",
	Run: func(cmd *cobra.Command, args []string) {
		dbm, err := tasks.NewDBManager(DBPath())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		todayList := dbm.List(time.Now())

		var data [][]string
		for _, task := range todayList {
			data = append(data, []string{
				task.Description,
				strconv.Itoa(*task.Effort),
				*task.EffortType,
			})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"description", "effort", "effort type"})

		for _, v := range data {
			table.Append(v)
		}

		table.Render()

	},
}

func init() {
	listCmd.AddCommand(todayCmd)
}
