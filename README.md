# Important points

- Used the `github.com/joho/godotenv` package in order to pull the PORT information from the .env to the current shell session
- Used `go mod vendor` to act as a local storage for modules, allowing Go to use these dependencies directly without needing to fetch them from remote repositories (kinnda like npm install)
- Used `go mod tidy` to clean up the go.mod and go.sum files by removing any unnecessary dependencies and adding any missing ones. This command is especially useful for maintaining a clean and optimized module configuration for your project
