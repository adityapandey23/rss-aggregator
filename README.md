# Important points

## Part-1 (Initial Setup)

- Used the `github.com/joho/godotenv` package in order to pull the PORT information from the .env to the current shell session
- Used `go mod vendor` to act as a local storage for modules, allowing Go to use these dependencies directly without needing to fetch them from remote repositories (kinnda like npm install)
- Used `go mod tidy` to clean up the go.mod and go.sum files by removing any unnecessary dependencies and adding any missing ones. This command is especially useful for maintaining a clean and optimized module configuration for your project

## Part-2 (Setting up the chi router)

- Installed the chi-router using `github.com/go-chi/chi` for routing and `github.com/go-chi/cors` for cors
- So, we have to first use some bits, then do `go mod vendor` which will bring the necessary code, and then we have to do `go mod tidy` to do some cleanups (yes it is confusing)
- Wrote some simple code for creating a router, createing a server and then make it listen to the port coming from the .env file
- Added some cors instructions which will help us to hit the server
