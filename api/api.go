package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Login string `json:"login"`
	PublicRepos int `json:"public_repos"`
	Email string `json:"email,omitempty"`
	FollowersNumber int `json:"followers"`
	CreatedAt time.Time `json:"created_at"` 
	Location string `json:"location"`
	Company string `json:"company,omitempty"`
}

type Event struct {
	EventType string `json:"type"`
	
	Repository struct{
		ID int64 `json:"id"`
		Name string `json:"name"`
	} `json:"repo"`
	
	Payload struct{
		Action string `json:"action"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		}`json:"commits"`
	}`json:"payload"`
}

func FetchEvents(username string) ([]string, error){
	eventsUrl := "https://api.github.com/users/" + username + "/events"
	var fetchedActivity []string
	var result []Event
	
	resp, err := http.Get(eventsUrl)
	if err != nil {
		return nil, err
	}

	if err := checkStatusCode(resp); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no activities found")
	}

	for _, event := range result {
		
		var output string

		switch event.EventType {
		case "PushEvent":
			pushCount := len(event.Payload.Commits)
			output = fmt.Sprintf("Pushed %d commits to %s", pushCount, event.Repository.Name) 
		
		case "CreateEvent":
			output = fmt.Sprintf("Created a %s in %s", event.Payload.RefType, event.Repository.Name)
		
		case "WatchEvent":
			output = fmt.Sprintf("Starred %s", event.Repository.Name)
		
		case "PullRequestEvent":
			output = fmt.Sprintf("Made a pull request in %s", event.Repository.Name)

		case "IssuesEvent":
			output = fmt.Sprintf("%s an issue in %s", event.Payload.Action, event.Repository.Name)

		default:
			output = fmt.Sprintf("%s in %s", event.EventType, event.Repository.Name)
		}

		fetchedActivity = append(fetchedActivity, output)
	}

	return fetchedActivity, nil
}

func FetchUserInfo(username string, locFlag bool, followersFlag bool) error{
	url := "https://api.github.com/users/" + username

	var user User

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := checkStatusCode(resp); err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return err
	}
	//year of creation account
	yearCreated := user.CreatedAt.Format("2006")

	//basic user's info
	fmt.Printf("User information:\nUsername: %s\nNumber of public repositories on user's Github: %d\n", user.Login, user.PublicRepos)
	if user.Email != "" {
		fmt.Println("User Email:", user.Email)
	}
	//company
	if user.Company != "" {
		fmt.Println("User is working in the", user.Company, "company")
	}
	
	//location
	if locFlag {
		fmt.Println("Location of user is", user.Location)
	}
	
	//followers
	if followersFlag {
		fmt.Println(user.FollowersNumber, "users have followed user's Github account")
	}
	
	fmt.Println("User's account was created in", yearCreated)

	return nil
}

func checkStatusCode(resp *http.Response) error{
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch user info. HTTP status code: %v", resp.StatusCode)
	}
	return nil
}