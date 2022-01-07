# GoDiscordBot Template
GoDiscordBot is a Discord bot template repository that helps you quickly get started builing a Discord bot in go.


## Development:
1. Compile and run the project.

    ```
    TOKEN=abc123 go run main.go
    ```

2. Alternatively, build and run the project inside of a container.

    ```
    docker build -t godiscordbot . && docker run -d --pull always --env TOKEN='abc123' godiscordbot
    ```