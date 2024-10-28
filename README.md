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

## Part-3 (Setting up a helper function)

- Making a helper function that will help in returning the JSON respose as per the request
- First, we'll marshal the data, if it fails we'll write status code 500 to the client otherwise, we'll first add a header to make the client side browser know that we are sending some JSON data, then we'll write the successful code as well as the data
- Added a handler_readiness function, which is just to check whether the server is health is good or not
- Then in the main file, we made a new router using chi (we made this because we want to have sub routes, well technically we don't but just for versioning and stuff), then we just mounted this new router (v1Router) to the original base route that we created
- Then we made an error response helper function, which will help us to know the error, (we used struct tag to mention the error for JSON), then we used the pervious function as well, then we finally made a new handler function and used it on the v1 router that we made previously
