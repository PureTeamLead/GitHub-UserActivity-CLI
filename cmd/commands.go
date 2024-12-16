package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var verbose string

func init() {
	rootCmd.AddCommand(getEvents)

	rootCmd.PersistentFlags().StringVarP(&verbose, "verbose", "v", "", "verbose output")
}

var getEvents = &cobra.Command{
	Use: "get",
	Short: "Command for fetching user activity",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching user activity...")
	},
}