package main

import (
	"fmt"
	"os"

	"gitlab.com/lcook/bugzport/command"
)

func main() {
	if _, err := command.RootCmd.ExecuteC(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
