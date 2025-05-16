package main

// https://www.aaronraff.dev/blog/how-to-write-a-lexer-in-go
// https://boyter.org/posts/faster-literal-string-matching-in-go/

// No LLMS were involved in the making of this program o7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Token int

const (
	T_EOF     Token = iota
	T_SHEBANG       // #!/bin/bash etc, ignore
	T_WORD          // command names, arguments
	T_NUMBER
	T_STRING
	T_PIPE
	T_ANDAND
	T_OROR
	T_DOLLAR
	T_SEMI
	T_COMMENT
	T_REDIRECT
	T_NEWLINE
)

var bash_tokens = []string{
	T_EOF:      "EOF",
	T_SHEBANG:  "#!", // discovery rule: ignore the rest of the line
	T_WORD:     "WORD",
	T_NUMBER:   "NUMBER",
	T_STRING:   "STRING",
	T_PIPE:     "|",
	T_ANDAND:   "&&",
	T_OROR:     "||",
	T_DOLLAR:   "$",
	T_SEMI:     ";",
	T_COMMENT:  "#",
	T_REDIRECT: ">",
	T_NEWLINE:  "\n",
}

var powershell_tokens = []string{
	T_NEWLINE: "\n",
	// todo
}

type FilledToken struct {
	typ Token
	val string
}

// sample bash command: ls -l | grep foo
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: powpiler <'compile', 'run'> <inputFile>")
		return
	}
	if args[1] == "compile" {
		if len(args) != 3 {
			fmt.Println("Usage: powpiler compile <inputFile>")
			return
		}

		f, err := os.Open(args[2])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		out, _ := compile(f)
		fmt.Printf("%s", out)

	} else if args[1] == "run" {
		fmt.Println("unimplemented") // TODO: run powershell and forward stdout/stderr
	}
}

// takes file and returns
func compile(f *os.File) (string, error) {
	reader := bufio.NewReader(f)
	lines, _ := FileToSlice(*reader, f)

	currChar := 0

	tokens := make([]FilledToken, 0)
	for _, v := range lines {
		for {
			idx := multiIndex(v, bash_tokens)
			tokens = append(tokens, FilledToken{typ: T_EOF, val: string(v[idx])})
			fmt.Printf("idx: %d\n", idx)
			currChar += idx + 1
		}
	}

	// print original input file
	for i := 0; i < len(lines); i++ {
		fmt.Printf("%s\n", lines[i])
	}

	return "", nil
}

// read a file to slice of strings
func FileToSlice(r bufio.Reader, f *os.File) ([]string, error) {
	outString := ""
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break // EOF
		}
		outString += string(line) + "\n"
	}
	lines := strings.Split(string(outString), "\n")
	return lines, nil
}

// read until a target substring string is found
// return index
// serious leetcode flashbacks
func multiIndex(s string, targets []string) int {
	for i := 0; i < len(s); i++ { // for each letter s[i]
		for j := 0; j < len(targets); j++ { // iterate over target targets[j]
			foundIdx := strings.Index(s[i:], targets[j])
			if foundIdx != -1 {
				return foundIdx + i
			}
		}
	}
	return -1
}
