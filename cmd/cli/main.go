package main

import (
	"flag"
	"fmt"
	"os"

	sfdl "github.com/lightDproject/SFDL"
)

func main() {
	parse := flag.Bool("parse", false, "Parse and validate SFDL file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: sfdl-cli [options] <file.sfdl>")
		flag.Usage()
		os.Exit(1)
	}

	file := flag.Arg(0)
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	config := sfdl.Config{Filename: file, Content: content}
	sfdlFile, err := sfdl.Parse(config)
	if err != nil {
		fmt.Printf("Parse errors: %v\n", err)
		os.Exit(1)
	}

	if *parse {
		fmt.Println("Parsed successfully")
		fmt.Printf("Attributes: %d\n", len(sfdlFile.SyntaxBody().Attributes))
		fmt.Printf("Blocks: %d\n", len(sfdlFile.SyntaxBody().Blocks))
	}
}
