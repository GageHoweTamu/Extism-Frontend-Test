package main

import (
	"bufio"
	"fmt"
	"os"
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
	T_NEWLINE: "`n",
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
	// outTokens := make([]FilledToken, 0)

	// fast file->memory reader; make slice of lines
	outLines := make([]string, 0)
	currentLine := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break // EOF
		}

		outLines = append(outLines, "")
		outLines[currentLine] += string(line) + "\n"
		currentLine++
	}

	// lexer
	char := 0
	line := 0
	buf := ""
	outTokens := make([]FilledToken, 0)
	for {
		switch outLines[char] {
		case '#':
			line++ // comment, ignore rest of line
		case '\n':
			outTokens = append(outTokens, FilledToken{T_NEWLINE, "\n"})
		default:
			// no recognized token
			buf += string(outString[char])
		}
		char++
	}

	// iterate over outLines
	for i := 0; i < len(outLines); i++ {
		fmt.Printf("%s", outLines[i])
	}

	return "", nil
}
