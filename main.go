package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/ogier/pflag"
)

var (
	user  string
	emoji string
)

func main() {
	pflag.Parse()
	// if user does not supply flags, print usage
	if pflag.NFlag() == 0 {
		printUsage()
	}
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	fmt.Println("")
	for _, u := range users {
		result := getUsers(u)
		color.Cyan(`Username:	%s`, result.Login)
		color.Blue(`Name:		%s`, result.Name)
		color.Green(`Email:		%s`, result.Email)
		color.HiMagenta(`Bio:		%s`, result.Bio)
		fmt.Println("")
	}

}

func init() {
	pflag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	pflag.PrintDefaults()
	os.Exit(1)
}

const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/"
)

// User struct for GitHub User
type User struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             string      `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func getUsers(name string) User {
	resp, err := http.Get(apiURL + userEndpoint + name)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var user User
	json.Unmarshal(body, &user)

	if (User{}) == user {

		color.Red("Sorry the GitHub user doesn't exist")
		os.Exit(1)
		return user
	}
	return user
}
