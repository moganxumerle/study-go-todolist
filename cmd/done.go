package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/moganxumerle/study-go-todolist/db"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ListAllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		for _, id := range ids {

			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Something went wrong when done the task %d. Error: %s", id, err)
			} else {
				fmt.Printf("The task %d - %s was done!", id, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
