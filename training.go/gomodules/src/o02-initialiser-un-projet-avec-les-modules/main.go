package main

import (
  log "github.com/sirupsen/logrus"
)

// commands
/**
$ go mod init training.go/helloworld

$ cat go.mod
module training.go/helloworld

go 1.14

# Add logrus library
$ go get

$ cat go.mod
module training.go/helloworld

go 1.14

require github.com/sirupsen/logrus v1.6.0

# Contains the cryptographic hashes of all the dependencies
$ cat go.sum
*/

/**
# Update dependency
$ go get github.com/sirupsen/logrus

# List all available dependencies
$ go list -u -m all

# Clean up unused dependencies
$ go mod tidy
*/

func main() {
    log.Info("Hello logrus")
}
