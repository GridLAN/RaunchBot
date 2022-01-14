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
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = os.Getenv("TOKEN")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

type RedditPost []struct {
	Kind string `json:"kind"`
	Data struct {
		After     interface{} `json:"after"`
		Dist      int         `json:"dist"`
		Modhash   string      `json:"modhash"`
		GeoFilter string      `json:"geo_filter"`
		Children  []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
				Subreddit                  string        `json:"subreddit"`
				Selftext                   string        `json:"selftext"`
				AuthorFullname             string        `json:"author_fullname"`
				Saved                      bool          `json:"saved"`
				ModReasonTitle             interface{}   `json:"mod_reason_title"`
				Gilded                     int           `json:"gilded"`
				Clicked                    bool          `json:"clicked"`
				Title                      string        `json:"title"`
				LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
				Hidden                     bool          `json:"hidden"`
				Pwls                       int           `json:"pwls"`
				LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
				Downs                      int           `json:"downs"`
				ThumbnailHeight            interface{}   `json:"thumbnail_height"`
				TopAwardedType             interface{}   `json:"top_awarded_type"`
				ParentWhitelistStatus      string        `json:"parent_whitelist_status"`
				HideScore                  bool          `json:"hide_score"`
				Name                       string        `json:"name"`
				Quarantine                 bool          `json:"quarantine"`
				LinkFlairTextColor         string        `json:"link_flair_text_color"`
				UpvoteRatio                float64       `json:"upvote_ratio"`
				AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
				SubredditType              string        `json:"subreddit_type"`
				Ups                        int           `json:"ups"`
				TotalAwardsReceived        int           `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        interface{}   `json:"thumbnail_width"`
				AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           interface{}   `json:"secure_media"`
				IsRedditMediaDomain   bool          `json:"is_reddit_media_domain"`
				IsMeta                bool          `json:"is_meta"`
				Category              interface{}   `json:"category"`
				SecureMediaEmbed      struct {
				} `json:"secure_media_embed"`
				LinkFlairText       interface{}   `json:"link_flair_text"`
				CanModPost          bool          `json:"can_mod_post"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				IsCreatedFromAdsUI  bool          `json:"is_created_from_ads_ui"`
				AuthorPremium       bool          `json:"author_premium"`
				Thumbnail           string        `json:"thumbnail"`
				Edited              bool          `json:"edited"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				Gildings            struct {
				} `json:"gildings"`
				ContentCategories   interface{} `json:"content_categories"`
				IsSelf              bool        `json:"is_self"`
				ModNote             interface{} `json:"mod_note"`
				CrosspostParentList []struct {
					ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
					Subreddit                  string        `json:"subreddit"`
					Selftext                   string        `json:"selftext"`
					UserReports                []interface{} `json:"user_reports"`
					Saved                      bool          `json:"saved"`
					ModReasonTitle             interface{}   `json:"mod_reason_title"`
					Gilded                     int           `json:"gilded"`
					Clicked                    bool          `json:"clicked"`
					Title                      string        `json:"title"`
					LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
					SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
					Hidden                     bool          `json:"hidden"`
					Pwls                       int           `json:"pwls"`
					LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
					Downs                      int           `json:"downs"`
					ThumbnailHeight            interface{}   `json:"thumbnail_height"`
					TopAwardedType             interface{}   `json:"top_awarded_type"`
					ParentWhitelistStatus      string        `json:"parent_whitelist_status"`
					HideScore                  bool          `json:"hide_score"`
					Name                       string        `json:"name"`
					Quarantine                 bool          `json:"quarantine"`
					LinkFlairTextColor         string        `json:"link_flair_text_color"`
					UpvoteRatio                float64       `json:"upvote_ratio"`
					AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
					SubredditType              string        `json:"subreddit_type"`
					Ups                        int           `json:"ups"`
					TotalAwardsReceived        int           `json:"total_awards_received"`
					MediaEmbed                 struct {
					} `json:"media_embed"`
					ThumbnailWidth        interface{} `json:"thumbnail_width"`
					AuthorFlairTemplateID interface{} `json:"author_flair_template_id"`
					IsOriginalContent     bool        `json:"is_original_content"`
					AuthorFullname        string      `json:"author_fullname"`
					SecureMedia           interface{} `json:"secure_media"`
					IsRedditMediaDomain   bool        `json:"is_reddit_media_domain"`
					IsMeta                bool        `json:"is_meta"`
					Category              interface{} `json:"category"`
					SecureMediaEmbed      struct {
					} `json:"secure_media_embed"`
					LinkFlairText       interface{}   `json:"link_flair_text"`
					CanModPost          bool          `json:"can_mod_post"`
					Score               int           `json:"score"`
					ApprovedBy          interface{}   `json:"approved_by"`
					IsCreatedFromAdsUI  bool          `json:"is_created_from_ads_ui"`
					AuthorPremium       bool          `json:"author_premium"`
					Thumbnail           string        `json:"thumbnail"`
					Edited              bool          `json:"edited"`
					AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
					AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
					Gildings            struct {
					} `json:"gildings"`
					ContentCategories        interface{}   `json:"content_categories"`
					IsSelf                   bool          `json:"is_self"`
					ModNote                  interface{}   `json:"mod_note"`
					Created                  float64       `json:"created"`
					LinkFlairType            string        `json:"link_flair_type"`
					Wls                      int           `json:"wls"`
					RemovedByCategory        interface{}   `json:"removed_by_category"`
					BannedBy                 interface{}   `json:"banned_by"`
					AuthorFlairType          string        `json:"author_flair_type"`
					Domain                   string        `json:"domain"`
					AllowLiveComments        bool          `json:"allow_live_comments"`
					SelftextHTML             string        `json:"selftext_html"`
					Likes                    interface{}   `json:"likes"`
					SuggestedSort            interface{}   `json:"suggested_sort"`
					BannedAtUtc              interface{}   `json:"banned_at_utc"`
					ViewCount                interface{}   `json:"view_count"`
					Archived                 bool          `json:"archived"`
					NoFollow                 bool          `json:"no_follow"`
					IsCrosspostable          bool          `json:"is_crosspostable"`
					Pinned                   bool          `json:"pinned"`
					Over18                   bool          `json:"over_18"`
					AllAwardings             []interface{} `json:"all_awardings"`
					Awarders                 []interface{} `json:"awarders"`
					MediaOnly                bool          `json:"media_only"`
					CanGild                  bool          `json:"can_gild"`
					Spoiler                  bool          `json:"spoiler"`
					Locked                   bool          `json:"locked"`
					AuthorFlairText          interface{}   `json:"author_flair_text"`
					TreatmentTags            []interface{} `json:"treatment_tags"`
					Visited                  bool          `json:"visited"`
					RemovedBy                interface{}   `json:"removed_by"`
					NumReports               interface{}   `json:"num_reports"`
					Distinguished            interface{}   `json:"distinguished"`
					SubredditID              string        `json:"subreddit_id"`
					AuthorIsBlocked          bool          `json:"author_is_blocked"`
					ModReasonBy              interface{}   `json:"mod_reason_by"`
					RemovalReason            interface{}   `json:"removal_reason"`
					LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
					ID                       string        `json:"id"`
					IsRobotIndexable         bool          `json:"is_robot_indexable"`
					ReportReasons            interface{}   `json:"report_reasons"`
					Author                   string        `json:"author"`
					DiscussionType           interface{}   `json:"discussion_type"`
					NumComments              int           `json:"num_comments"`
					SendReplies              bool          `json:"send_replies"`
					Media                    interface{}   `json:"media"`
					ContestMode              bool          `json:"contest_mode"`
					AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
					AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
					Permalink                string        `json:"permalink"`
					WhitelistStatus          string        `json:"whitelist_status"`
					Stickied                 bool          `json:"stickied"`
					URL                      string        `json:"url"`
					SubredditSubscribers     int           `json:"subreddit_subscribers"`
					CreatedUtc               float64       `json:"created_utc"`
					NumCrossposts            int           `json:"num_crossposts"`
					ModReports               []interface{} `json:"mod_reports"`
					IsVideo                  bool          `json:"is_video"`
				} `json:"crosspost_parent_list"`
				Created                  float64       `json:"created"`
				LinkFlairType            string        `json:"link_flair_type"`
				Wls                      int           `json:"wls"`
				RemovedByCategory        interface{}   `json:"removed_by_category"`
				BannedBy                 interface{}   `json:"banned_by"`
				AuthorFlairType          string        `json:"author_flair_type"`
				Domain                   string        `json:"domain"`
				AllowLiveComments        bool          `json:"allow_live_comments"`
				SelftextHTML             interface{}   `json:"selftext_html"`
				Likes                    interface{}   `json:"likes"`
				SuggestedSort            interface{}   `json:"suggested_sort"`
				BannedAtUtc              interface{}   `json:"banned_at_utc"`
				URLOverriddenByDest      string        `json:"url_overridden_by_dest"`
				ViewCount                interface{}   `json:"view_count"`
				Archived                 bool          `json:"archived"`
				NoFollow                 bool          `json:"no_follow"`
				IsCrosspostable          bool          `json:"is_crosspostable"`
				Pinned                   bool          `json:"pinned"`
				Over18                   bool          `json:"over_18"`
				AllAwardings             []interface{} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          interface{}   `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				AuthorIsBlocked          bool          `json:"author_is_blocked"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				NumDuplicates            int           `json:"num_duplicates"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				Media                    interface{}   `json:"media"`
				ContestMode              bool          `json:"contest_mode"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				CrosspostParent          string        `json:"crosspost_parent"`
				AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				WhitelistStatus          string        `json:"whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				ModReports               []interface{} `json:"mod_reports"`
				IsVideo                  bool          `json:"is_video"`
			} `json:"data"`
		} `json:"children"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

type Subreddit struct {
	Kind string `json:"kind"`
	Data struct {
		UserFlairBackgroundColor       interface{}   `json:"user_flair_background_color"`
		SubmitTextHTML                 string        `json:"submit_text_html"`
		RestrictPosting                bool          `json:"restrict_posting"`
		UserIsBanned                   interface{}   `json:"user_is_banned"`
		FreeFormReports                bool          `json:"free_form_reports"`
		WikiEnabled                    bool          `json:"wiki_enabled"`
		UserIsMuted                    interface{}   `json:"user_is_muted"`
		UserCanFlairInSr               interface{}   `json:"user_can_flair_in_sr"`
		DisplayName                    string        `json:"display_name"`
		HeaderImg                      string        `json:"header_img"`
		Title                          string        `json:"title"`
		AllowGalleries                 bool          `json:"allow_galleries"`
		IconSize                       []int         `json:"icon_size"`
		PrimaryColor                   string        `json:"primary_color"`
		ActiveUserCount                int           `json:"active_user_count"`
		IconImg                        string        `json:"icon_img"`
		DisplayNamePrefixed            string        `json:"display_name_prefixed"`
		AccountsActive                 int           `json:"accounts_active"`
		PublicTraffic                  bool          `json:"public_traffic"`
		Subscribers                    int           `json:"subscribers"`
		UserFlairRichtext              []interface{} `json:"user_flair_richtext"`
		VideostreamLinksCount          int           `json:"videostream_links_count"`
		Name                           string        `json:"name"`
		Quarantine                     bool          `json:"quarantine"`
		HideAds                        bool          `json:"hide_ads"`
		PredictionLeaderboardEntryType string        `json:"prediction_leaderboard_entry_type"`
		EmojisEnabled                  bool          `json:"emojis_enabled"`
		AdvertiserCategory             string        `json:"advertiser_category"`
		PublicDescription              string        `json:"public_description"`
		CommentScoreHideMins           int           `json:"comment_score_hide_mins"`
		AllowPredictions               bool          `json:"allow_predictions"`
		UserHasFavorited               interface{}   `json:"user_has_favorited"`
		UserFlairTemplateID            interface{}   `json:"user_flair_template_id"`
		CommunityIcon                  string        `json:"community_icon"`
		BannerBackgroundImage          string        `json:"banner_background_image"`
		OriginalContentTagEnabled      bool          `json:"original_content_tag_enabled"`
		CommunityReviewed              bool          `json:"community_reviewed"`
		SubmitText                     string        `json:"submit_text"`
		DescriptionHTML                string        `json:"description_html"`
		SpoilersEnabled                bool          `json:"spoilers_enabled"`
		HeaderTitle                    string        `json:"header_title"`
		HeaderSize                     []int         `json:"header_size"`
		UserFlairPosition              string        `json:"user_flair_position"`
		AllOriginalContent             bool          `json:"all_original_content"`
		HasMenuWidget                  bool          `json:"has_menu_widget"`
		IsEnrolledInNewModmail         interface{}   `json:"is_enrolled_in_new_modmail"`
		KeyColor                       string        `json:"key_color"`
		CanAssignUserFlair             bool          `json:"can_assign_user_flair"`
		Created                        int           `json:"created"`
		Wls                            int           `json:"wls"`
		ShowMediaPreview               bool          `json:"show_media_preview"`
		SubmissionType                 string        `json:"submission_type"`
		UserIsSubscriber               interface{}   `json:"user_is_subscriber"`
		DisableContributorRequests     bool          `json:"disable_contributor_requests"`
		AllowVideogifs                 bool          `json:"allow_videogifs"`
		ShouldArchivePosts             bool          `json:"should_archive_posts"`
		UserFlairType                  string        `json:"user_flair_type"`
		AllowPolls                     bool          `json:"allow_polls"`
		CollapseDeletedComments        bool          `json:"collapse_deleted_comments"`
		EmojisCustomSize               interface{}   `json:"emojis_custom_size"`
		PublicDescriptionHTML          string        `json:"public_description_html"`
		AllowVideos                    bool          `json:"allow_videos"`
		IsCrosspostableSubreddit       bool          `json:"is_crosspostable_subreddit"`
		NotificationLevel              interface{}   `json:"notification_level"`
		CanAssignLinkFlair             bool          `json:"can_assign_link_flair"`
		AccountsActiveIsFuzzed         bool          `json:"accounts_active_is_fuzzed"`
		AllowPredictionContributors    bool          `json:"allow_prediction_contributors"`
		SubmitTextLabel                string        `json:"submit_text_label"`
		LinkFlairPosition              string        `json:"link_flair_position"`
		UserSrFlairEnabled             interface{}   `json:"user_sr_flair_enabled"`
		UserFlairEnabledInSr           bool          `json:"user_flair_enabled_in_sr"`
		AllowDiscovery                 bool          `json:"allow_discovery"`
		AcceptFollowers                bool          `json:"accept_followers"`
		UserSrThemeEnabled             bool          `json:"user_sr_theme_enabled"`
		LinkFlairEnabled               bool          `json:"link_flair_enabled"`
		SubredditType                  string        `json:"subreddit_type"`
		SuggestedCommentSort           interface{}   `json:"suggested_comment_sort"`
		BannerImg                      string        `json:"banner_img"`
		UserFlairText                  interface{}   `json:"user_flair_text"`
		BannerBackgroundColor          string        `json:"banner_background_color"`
		ShowMedia                      bool          `json:"show_media"`
		ID                             string        `json:"id"`
		UserIsModerator                interface{}   `json:"user_is_moderator"`
		Over18                         bool          `json:"over18"`
		Description                    string        `json:"description"`
		SubmitLinkLabel                string        `json:"submit_link_label"`
		UserFlairTextColor             interface{}   `json:"user_flair_text_color"`
		RestrictCommenting             bool          `json:"restrict_commenting"`
		UserFlairCSSClass              interface{}   `json:"user_flair_css_class"`
		AllowImages                    bool          `json:"allow_images"`
		Lang                           string        `json:"lang"`
		WhitelistStatus                string        `json:"whitelist_status"`
		URL                            string        `json:"url"`
		CreatedUtc                     int           `json:"created_utc"`
		BannerSize                     []int         `json:"banner_size"`
		MobileBannerImage              string        `json:"mobile_banner_image"`
		UserIsContributor              interface{}   `json:"user_is_contributor"`
		AllowPredictionsTournament     bool          `json:"allow_predictions_tournament"`
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
			Name:        "reddit",
			Description: "random post from a random subreddit",
		},
		{
			Name:        "reddit-add",
			Description: "add a subreddit to the list of available subreddits",
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
			Name:        "reddit-list",
			Description: "list of available subreddits",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "reddit-remove",
			Description: "remove a subreddit to the list of available subreddits",
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
			Name:        "subreddit",
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
		"reddit": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var msg string
			// If the []subreddits is empty, it will return an error
			if len(Subreddits) == 0 {
				msg = "There are no subreddits in the list to choose from."
			} else {

				// Seed the random number generator & get a random index from the slice
				rand.Seed(time.Now().UnixNano())
				randIndex := rand.Intn(len(Subreddits))

				// Query the API for a random post from a random subreddit
				randomRedditPost := RedditPost{}
				getJson("https://reddit.com/r/"+Subreddits[randIndex]+"/random.json", &randomRedditPost)

				msg = fmt.Sprintf(randomRedditPost[0].Data.Children[0].Data.Title) + " " + randomRedditPost[0].Data.Children[0].Data.URL
			}
			// Respond with the post's title and URL
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})
		},
		"reddit-add": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Query user for subreddit name
			subreddit := i.ApplicationCommandData().Options[0].StringValue()

			// Setup the response data
			var msg string

			subredditCheck := Subreddit{}
			getJson("https://reddit.com/r/"+subreddit+"/about.json", &subredditCheck)

			// If the subreddit exists, add it to the slice
			if subredditCheck.Data.URL == "" {
				msg = "The subreddit " + subreddit + " was not found. Try again."
			} else {
				Subreddits = append(Subreddits, subreddit)
				msg = subreddit + " has been added to the list."
			}

			// Respond with the message
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})

		},
		"reddit-list": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			msg := "The following subreddits are available:\n" + "```" + strings.Join(Subreddits, "\n") + "```"

			// Respond with the message
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})
		},
		"reddit-remove": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		"subreddit": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Query user for subreddit name
			subreddit := i.ApplicationCommandData().Options[0].StringValue()

			var msg string
			// If the []subreddits is empty, it will return an error
			if len(Subreddits) == 0 {
				msg = "There are no subreddits in the list to choose from."
			}
			// If subreddit is not in []subreddits, return error
			if !contains(Subreddits, subreddit) {
				msg = subreddit + " is not in the list."
			} else {

				// Query the API for a random post from a random subreddit
				randomRedditPost := RedditPost{}
				getJson("https://reddit.com/r/"+subreddit+"/random.json", &randomRedditPost)

				msg = fmt.Sprintf(randomRedditPost[0].Data.Children[0].Data.Title) + " " + randomRedditPost[0].Data.Children[0].Data.URL
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
	log.Println("Gracefully shutdowning")
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
