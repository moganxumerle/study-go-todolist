package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/moganxumerle/study-go-todolist/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to you task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong to add your task to list: ", err)
			os.Exit(1)
		}

		fmt.Printf("Added \"%s\" to you task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
