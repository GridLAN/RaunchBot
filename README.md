# RaunchBot
RaunchBot brings you the latest Raunchy content straight to your favorite Discord server. 

## Development:
1. Compile and run the project.

    ```
    TOKEN=abc123 go run main.go
    ```

2. Alternatively, build and run the project inside of a container.

    ```
    docker build -t raunchbot . && docker run -d --pull always --env TOKEN='abc123' raunchbot
    ```
