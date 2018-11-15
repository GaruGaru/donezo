package cmd

import (
	"github.com/GaruGaru/done/tasks"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
	"path"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo is a cli helper for devs that needs to work instead of getting micromanaged",
}

func DBPath() string {
	usr, err := user.Current()
	if err != nil {
		path.Join("/tmp/", ".todo.db")
	}
	return path.Join(usr.HomeDir, ".todo.db")
}

func Repository() tasks.Repository {
	db, err := tasks.NewDBRepository(DBPath())
	if err != nil {
		panic(err)
	}
	return db
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
