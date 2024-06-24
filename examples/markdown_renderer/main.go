package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/JoshPattman/obtext"
	"github.com/JoshPattman/obtext/markup"
)

func main() {
	// Specify the command line args and parse them
	var inputFileName string
	var outputFileName string
	var secondaryOutputFileName string
	var prettyPrint bool
	var generateHtml bool

	flag.StringVar(&inputFileName, "i", "", "The input file (.obt) to read")
	flag.StringVar(&outputFileName, "o", "", "The output file (.md) to write")
	flag.BoolVar(&prettyPrint, "p", false, "Pretty print the parsed obt ast to output")
	flag.BoolVar(&generateHtml, "h", false, "Generate an html file as well as a markdown file")

	flag.Parse()

	// Ensure we have an input file and output file
	if inputFileName == "" {
		fmt.Println("No input file specified")
		os.Exit(1)
	}
	if outputFileName == "" {
		outputFileName = strings.TrimSuffix(inputFileName, ".obt") + ".md"
	}

	secondaryOutputFileName = strings.TrimSuffix(outputFileName, ".md") + ".html"

	// Read the content of the input file
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Failed to open input file:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// Try to parse the input file into an AST (the syntax parsing step)
	ast, err := obtext.ParseSynReader(inputFile)
	if err != nil {
		fmt.Println("Failed to parse input file:", err)
		os.Exit(1)
	}

	// Pretty print the AST if requested
	if prettyPrint {
		fmt.Println(obtext.FormatSynWithAnsiiColors(ast))
	}

	// Try to parse the syntax tree into a semantic tree (the semantics parsing step)
	st, err := obtext.ParseSem(ast, markup.Semantics)
	if err != nil {
		fmt.Println("Failed to process semantics:", err)
		os.Exit(1)
	}

	// Create the output file
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Generate the markdown from the semantic tree
	md := markup.RenderMarkdown(st)
	fmt.Fprint(outputFile, md)

	// Success!
	fmt.Println("Successfully wrote markdown to", outputFileName)

	// Generate the html file if requested
	if generateHtml {
		// Create the output file
		htmlOutputFile, err := os.Create(secondaryOutputFileName)
		if err != nil {
			fmt.Println("Failed to create output file:", err)
			os.Exit(1)
		}
		defer htmlOutputFile.Close()

		// Generate the html from the semantic tree
		html := markup.RenderHTML(st, "")
		fmt.Fprint(htmlOutputFile, html)

		// Success!
		fmt.Println("Successfully wrote html to", secondaryOutputFileName)
	}
}
