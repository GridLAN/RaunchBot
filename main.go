package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID  = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken = os.Getenv("TOKEN")
)

type RedditPost []struct {
	Kind string `json:"kind"`
	Data struct {
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				Subreddit string `json:"subreddit"`
				Title     string `json:"title,omitempty"`
				URL       string `json:"url,omitempty"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type Subreddit struct {
	Kind string `json:"kind"`
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
}

var Subreddits = []string{}

func getJson(url string, target interface{}) error {
	// Create a new HTTP client with a timeout of 10 seconds
	var myClient = http.Client{Timeout: 10 * time.Second}

	// Build a GET request to the meal API endpoint
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add the required headers to the request
	req.Header.Set("User-Agent", "RaunchBot")

	// Send the request and store the response
	r, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// Close the response body when done
	defer r.Body.Close()

	// Read the response body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON response into the target interface
	err = json.Unmarshal(b, &target)
	if err != nil {
		log.Fatal(err)
	}

	return json.Unmarshal(b, target)
}

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	if BotToken == "" {
		log.Fatal("Token cannot be empty")
	}

	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "random",
			Description: "random post from a random subreddit in the list",
		},
		{
			Name:        "add",
			Description: "add a subreddit to the list",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "subreddit",
					Description: "enter subreddit name",
					Required:    true,
				},
			},
		},
		{
			Name:        "remove",
			Description: "remove a subreddit from the list",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "subreddit",
					Description: "enter subreddit name",
					Required:    true,
				},
			},
		},
		{
			Name:        "list",
			Description: "lists of available subreddits in the list",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "sub",
			Description: "random post from a specific subreddit",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "subreddit",
					Description: "enter subreddit name",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"random": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var msg string
			// If the []subreddits is empty, it will return an error
			if len(Subreddits) == 0 {
				msg = "There are no subreddits on any list."
			} else {

				// Seed the random number generator & get a random index from the slice
				rand.Seed(time.Now().UnixNano())
				randIndex := rand.Intn(len(Subreddits))

				// print subreddit
				fmt.Println(Subreddits[randIndex])

				// Query the API for a random post from a random subreddit
				randomRedditPost := RedditPost{}
				getJson("https://reddit.com/r/"+Subreddits[randIndex]+"/random.json?obey_over18=true", &randomRedditPost)

				msg = randomRedditPost[0].Data.Children[0].Data.Title + "\n`r/" + randomRedditPost[0].Data.Children[0].Data.Subreddit + "`\n" + randomRedditPost[0].Data.Children[0].Data.URL
			}
			// Respond with the post's title and URL
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})
		},
		"add": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Query user for subreddit name
			subreddit := i.ApplicationCommandData().Options[0].StringValue()

			// Setup the response data
			var msg string

			// if the subreddit is already contained in the slice, do nothing
			if contains(Subreddits, subreddit) {
				msg = "The subreddit " + subreddit + " is already on the list."
			} else {
				subredditCheck := Subreddit{}
				getJson("https://reddit.com/r/"+subreddit+"/about.json", &subredditCheck)
				// If the subreddit exists, add it to the slice
				if subredditCheck.Data.URL == "" {
					msg = "The subreddit " + subreddit + " was not found. Try again."
				} else {
					Subreddits = append(Subreddits, subreddit)
					msg = subreddit + " has been added to the list."
				}
			}
			// Respond with the message
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})

		},
		"list": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Setup the response data
			var msg string

			// If the []subreddits is empty, it will return an error
			if len(Subreddits) == 0 {
				msg = "There are no subreddits any list."
			} else {
				msg = "The following subreddits are available:\n" + "```\n" + strings.Join(Subreddits, "\n") + "\n```"
			}

			// Respond with the message
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})
		},
		"remove": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Query user for subreddit name
			subreddit := i.ApplicationCommandData().Options[0].StringValue()

			// Setup the response data
			var msg string

			// If subreddit is not in []subreddits, return error
			if !contains(Subreddits, subreddit) {
				msg = subreddit + " is not in the list."
			} else {
				// Remove subreddit from []subreddits
				Subreddits = remove(Subreddits, subreddit)
				msg = subreddit + " has been removed from the list."
			}

			// Respond with the message
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})
		},
		"sub": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Query user for subreddit name
			subreddit := i.ApplicationCommandData().Options[0].StringValue()

			// Setup the response data
			var msg string

			subredditCheck := Subreddit{}
			getJson("https://reddit.com/r/"+subreddit+"/about.json", &subredditCheck)

			// If the subreddit exists, get a random post from it
			if subredditCheck.Data.URL == "" {
				msg = "The subreddit " + subreddit + " was not found. Try again."
			} else {
				// Query the API for a random post from a random subreddit
				randomRedditPost := RedditPost{}
				getJson("https://reddit.com/r/"+subreddit+"/random.json", &randomRedditPost)

				msg = randomRedditPost[0].Data.Children[0].Data.Title + "\n`r/" + randomRedditPost[0].Data.Children[0].Data.Subreddit + "`\n" + randomRedditPost[0].Data.Children[0].Data.URL
			}

			// Respond with the post's title and URL
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})
		},
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	defer s.Close()
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Gracefully shutdowning; Cleaning up commands")

	for _, v := range commands {
		s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.Name)
	}
}

// helper functions
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
