package main

import (
	"fmt"
	"os"
)

type Token int

const (
	T_EOF  Token = iota
	T_WORD       // command names, arguments
	T_NUMBER
	T_STRING
	T_PIPE
	T_ANDAND
	T_OROR
	T_DOLLAR
	T_SEMI
	T_COMMENT
)

// sample bash command: ls -l | grep foo
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: powpiler <-compile OR -run> <input>")
		return
	}
	if args[1] == "-compile" {
		if len(args) != 3 {
			fmt.Println("Usage: powpiler -compile <input>")
			return
		}
		input := args[2]
		fmt.Printf("%s", input) // TODO: output compiled powershell
	} else if args[1] == "-run" {
		fmt.Println("unimplemented") // TODO: run powershell and forward stdout/stderr
	}
}
