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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "print tasks list",
}

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "list tasks for the current day",
	Run: func(cmd *cobra.Command, args []string) {
		repository := Repository()
		todayList := repository.List(time.Now())
		prettyPrintTasks(todayList)
	},
}

var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Short: "list tasks for yesterday",
	Run: func(cmd *cobra.Command, args []string) {
		repository := Repository()
		yesterdayTime := time.Now().Add(-24 * time.Hour)
		prettyPrintTasks(repository.List(yesterdayTime))
	},
}

var dayCmd = &cobra.Command{
	Use:   "day [dd-mm-yyyy]",
	Short: "list tasks for a specific day",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("no time provided")
			return
		}
		repository := Repository()
		dayTime, err := time.Parse("02-01-2006", args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		prettyPrintTasks(repository.List(dayTime))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(todayCmd)
	listCmd.AddCommand(yesterdayCmd)
	listCmd.AddCommand(dayCmd)
}

func prettyPrintTasks(tasks []tasks.Task) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"description", "effort", "effort type"})

	var data [][]string
	for _, task := range tasks {
		data = append(data, []string{
			task.Description,
			strconv.Itoa(*task.Effort),
			*task.EffortType,
		})
	}

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
