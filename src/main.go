package main

import (
	"fmt"
	"monkeylang/src/repl"
	"os"
	"os/user"
)

func main() {

	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Monkey v0.0.0\n", usr.Username)
	repl.Start(os.Stdin, os.Stdout)
}
