package main

import (
	//	"bufio"
	"fmt"
	"log"
	"os"
	"sqlriver/pkg/lexer"
	"sqlriver/pkg/utils"
)

func runFile(path string) {

	// does file exist
	_, err := os.Stat(path)
	utils.Check(err)
	// try opening
	dat, err := os.ReadFile(path)
	utils.Check(err)

	source := string(dat)
	log.Println("Read value: ")
	log.Println(source)
	run(source)

}

func run(source string) {
	lexer := lexer.NewLexer(source)
	tokens := lexer.Lex()
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func main() {
	nArgs := len(os.Args)

	if nArgs > 2 {
		fmt.Println("Usage: sqlriver [input_file]")
		os.Exit(64)
	} else if nArgs == 2 {
		fp := os.Args[1]
		log.Println("Running on file at ", fp)
		runFile(fp)
	} else {
		//runStdIn()i
		fmt.Println("Running from stdin")
	}

}
