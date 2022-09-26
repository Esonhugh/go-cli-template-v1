package main

import (
	"fast_cli_template/cmd"
	_ "fast_cli_template/cmd/version" // import sub command as module
)

func init() {
}

func main() {
	cmd.Execute()
}
