package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "github-activity",
	Short: "Github CLI app fetches user's events",
	Long: `GitHub CLI application primary use is for fetching user's data 
	about his account events such as posting different repos or
	making commits and much more. 
	Enjoy the CLI app!`,

	Run: func(cmd *cobra.Command, args []string) {
		if AuthorPersistentFlag {
			fmt.Println("Author of the project is PureTeamLead")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}