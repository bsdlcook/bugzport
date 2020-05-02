package main

import (
	"fmt"
	"gitlab.com/lcook/bugzport/command"
	"os"
)

func main() {
	if _, err := command.RootCmd.ExecuteC(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
