package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sqlriver/pkg/lexer"
)

func runFile(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	source := string(dat)
	log.Print("Read input from file:\n", source)
	log.Print(len(source))
	if err := run(source); err != nil {
		log.Fatal(err)
	}

}

func runPrompt() {
	// TODO: Add meta-commands for help, quit etc.
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print(">")
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			log.Fatal("Error reading from stdin: ", err)
		}
		fmt.Print("Got input: ", text)
		if err := run(text); err != nil {
			fmt.Println(err)
		}
	}
}

func run(source string) error {
	lexer := lexer.NewLexer(source)
	tokens, err := lexer.Lex()
	if err != nil {
		return err
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}

func main() {
	nArgs := len(os.Args)

	if nArgs > 2 {
		// TODO add better usage prompts
		fmt.Println("Usage: sqlriver [input_file]")
		os.Exit(64)
	} else if nArgs == 2 {
		fp := os.Args[1]
		log.Println("Running on file at ", fp)
		runFile(fp)
	} else {
		fmt.Println("Running in REPL prompt")
		runPrompt()
	}

}
