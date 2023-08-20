# quick-go-demo

## Setup

1. run `docker compose up` to launch the postgres service and load up the data that is stored in `sql` folder.  Will take a few minutes since the script is not optimized at all.

1. run `go mod init <package-name>`

1. Create a `main.go` file

1. Create a repository package that connects to the postgres database and returns a list of cards based on likeness search on the card name

1. Import `github.com/go-chi/chi` in `main.go` to begin creating the router.
    * I'm choosing `chi` framework because it adds very little extra to the go std lib to help show what go can do.

## Next steps to look at

1. Look at `gRPC`
1. Look at `https://github.com/mitranim/gow` for file watching and rebuilding on file changes
