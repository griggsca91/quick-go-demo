# quick-go-demo

## Setup

1. run `docker compose up` to launch the postgres service and load up the data that is stored in `sql` folder.  Will take a few minutes since the script is not optimized at all.

1. run `go mod init <package-name>`

1. Create a `main.go` file

1. Import `github.com/go-chi/chi/v5` in `main.go` to begin creating the router.
    * I'm choosing `chi` framework because it adds very little extra to the go std lib to help show what go can do.

1. Write a "Hello World" route that returns a simple json. healthcheck route on `chi.NewRouter`

1. Launch server on port `:3333` with `http.ListenAndServe` and pass in the `chi.Router`

1. `curl localhost:3333/healthcheck` to see

1. Run `go mod tidy` to automatically import dependencies

1. Run `gofmt -w .` to automatically format your code

1. add a secure middleware that prevents the user from accessing the api unless they put a password query param in the url

1. use `localhost:3333/healthcheck\?password=test` to access the system securely now

1. Create a route that is used to return a list of cards with similar names

1. `curl localhost:3333/cards\?password`

1. Create an interface for the respository we want to use to get these cards

1. Create a repository package that connects to the postgres database and returns a list of cards based on likeness search on the card name

1. Import `github.com/jackc/pgx/v5` and create a pool to pass in to the repository impl


## Next steps to look at

1. Look at `gRPC`
1. Look at `https://github.com/mitranim/gow` for file watching and rebuilding on file changes
1. Look at `https://github.com/mvdan/gofumpt`
1. Code review advice for reviewing go code - `https://github.com/golang/go/wiki/CodeReviewComments#receiver-type`