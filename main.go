package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ammartinez008/interpreter-golang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! this is Monkey Lang Programming Language! \n", user.Username)
	fmt.Printf("start using\n")
	repl.Start(os.Stdin, os.Stdout)
}
