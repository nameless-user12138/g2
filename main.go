package main

import (
	"fmt"
	"os"
	"os/user"
	"g2/repl"
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