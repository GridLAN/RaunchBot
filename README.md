# RedditBot
Get RedditBot in your discord [here](https://discord.com/api/oauth2/authorize?client_id=931745853331374170&permissions=534723951680&scope=bot%20applications.commands).

Bot Commands:
```
/random - random post from random subreddit
/list - show the list of subreddits
/add - add subreddit to the list
/remove - remove subreddit from the list
/sub <subreddit> - random post from <subreddit>
```

## Development:
1. Compile and run the project.

    ```
    TOKEN=abc123 go run main.go
    ```

2. Alternatively, build and run the project inside of a container.

    ```
    docker build -t redditbot . && docker run -d --env TOKEN='abc123' redditbot
    ```
