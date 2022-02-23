package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/moganxumerle/study-go-todolist/cmd"
	"github.com/moganxumerle/study-go-todolist/db"
)

func main() {
	home, _ := homedir.Dir()
	fmt.Println(home)
	dbPath := filepath.Join(home, "task.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
