package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sfdl <file.sfdl>")
		os.Exit(1)
	}

	file := os.Args[1]
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	hclFile, diags := hclsyntax.ParseConfig(content, file, hcl.InitialPos)
	if diags.HasErrors() {
		fmt.Printf("Parse errors: %v\n", diags)
		os.Exit(1)
	}

	fmt.Println("Parsed HCL AST successfully")
	fmt.Printf("Body: %+v\n", hclFile.Body)
}
