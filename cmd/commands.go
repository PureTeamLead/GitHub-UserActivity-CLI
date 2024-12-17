package cmd

import (
	"fmt"
	"gitHub_fetch_cli/api"

	"github.com/spf13/cobra"
)

var locationFlag bool
var followersFlag bool
var AuthorPersistentFlag bool

func init() {
	rootCmd.AddCommand(getEvents, getUserInfo)
	getUserInfo.Flags().BoolVarP(&locationFlag, "location", "l", false, "Flag is used for fetching user's location")
	getUserInfo.Flags().BoolVarP(&followersFlag, "followers", "f", false, "Flag is used for fetching number of user's followers")
	rootCmd.PersistentFlags().BoolVarP(&AuthorPersistentFlag, "author", "a", false, "Flag is used for information about author of the project")
}

var getEvents = &cobra.Command{
	Use: "get_events [github-username]",
	Short: "Command for fetching user activity",
	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error{
		
		username := args[0]

		fmt.Println("Fetching user activity...")

		eventsInfo, err := api.FetchEvents(username)
		if err != nil {
			return err
		}

		for _, eventInfo := range eventsInfo {
			fmt.Println(eventInfo)
		}
		
		return nil
	},
}

var getUserInfo = &cobra.Command{
	Use: "get_info [github-username]",
	Short: "Command for fetching user's info",
	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		if err := api.FetchUserInfo(username, locationFlag, followersFlag); err != nil {
			return err
		}

		return nil
	},
}