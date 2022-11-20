# CKGameRollBot
A Discord Bot, written in Go, to help my friends and I decide what game to play when none of us know what we want to do.

## Managing Dependencies
### To add dependencies:
There are two ways to do this.
1. Run `go get [URL]` (i.e. `go get rsc.io/quote`)
2. Use the code as if it was installed and then run `go mod tidy`
### To remove dependencies:
1. Remove all the imports and code related to the package
2. Run `go mod tidy -v` command to remove any unused package