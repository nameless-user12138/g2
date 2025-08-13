package main

import (
	"fmt"
	"g2/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s.\n", user.Name)
	fmt.Printf("Welcome to use monkey script language.\n")
	repl.Start(os.Stdin, os.Stdout)
}
