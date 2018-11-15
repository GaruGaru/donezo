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
	Use:   "donezo",
	Short: "donezo is a cli helper for devs that needs to work instead of getting micromanaged",
}

func DBPath() string {
	usr, err := user.Current()
	if err != nil {
		path.Join("/tmp/", ".donezo.db")
	}
	return path.Join(usr.HomeDir, ".donezo.db")
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
