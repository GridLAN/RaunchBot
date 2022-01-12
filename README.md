# RaunchBot
RaunchBot brings you the latest Raunchy content straight to your favorite Discord server. 

Get RanchBot in your discord [here](https://discord.com/api/oauth2/authorize?client_id=929197630981033995&permissions=534723951680&scope=bot%20applications.commands).

Bot Commands:
```
/reddit - random post from random subreddit
/reddit add - add subreddit to []subreddits{}
/reddit remove - remove subreddit from []subreddits{}
/subreddit <subreddit> - random post from <subreddit>
```

## Development:
1. Compile and run the project.

    ```
    TOKEN=abc123 go run main.go
    ```

2. Alternatively, build and run the project inside of a container.

    ```
    docker build -t raunchbot . && docker run -d --pull always --env TOKEN='abc123' raunchbot
    ```
