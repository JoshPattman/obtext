package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/JoshPattman/obtext"
)

func main() {
	// Specify the command line args and parse them
	var inputFileName string
	var outputFileName string
	var prettyPrint bool

	flag.StringVar(&inputFileName, "i", "", "The input file (.obt) to read")
	flag.StringVar(&outputFileName, "o", "", "The output file (.md) to write")
	flag.BoolVar(&prettyPrint, "p", false, "Pretty print the parsed obt ast to output")

	flag.Parse()

	// Ensure we have an input file and output file
	if inputFileName == "" {
		fmt.Println("No input file specified")
		os.Exit(1)
	}
	if outputFileName == "" {
		outputFileName = strings.TrimSuffix(inputFileName, ".obt") + ".md"
	}

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
	st, err := obtext.ParseSem(ast, obtext.DefaultMarkupSemantics)
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
	md := generateMarkdown(st)
	fmt.Fprint(outputFile, md)

	// Success!
	fmt.Println("Successfully wrote markdown to", outputFileName)
}

// generateMarkdown takes a semantic tree and generates a markdown string from it.
// Rendering like this is deliberately left out of the obtext package, as it is intended to be done by the user of the package.
// For example, I wrote a renderer for my personal blog that uses templ to render the semantic tree into HTML directly.
func generateMarkdown(t obtext.SemNode) string {
	switch t := t.(type) {
	case *obtext.ContentBlockSemNode:
		out := ""
		for _, e := range t.Elements {
			out += generateMarkdown(e)
		}
		return out
	case *obtext.TextSemNode:
		return t.Text
	case *obtext.DocSemNode:
		return generateMarkdown(t.Content)
	case *obtext.H1SemNode:
		return "\n# " + generateMarkdown(t.Content) + "\n"
	case *obtext.H2SemNode:
		return "\n## " + generateMarkdown(t.Content) + "\n"
	case *obtext.PSemNode:
		return "\n" + generateMarkdown(t.Content) + "\n"
	case *obtext.BoldSemNode:
		return "**" + generateMarkdown(t.Content) + "**"
	case *obtext.ItalicSemNode:
		return "*" + generateMarkdown(t.Content) + "*"
	case *obtext.ImageSemNode:
		return fmt.Sprintf("\n![%s](%s)\n", generateMarkdown(t.CaptionContent), t.Link)
	case *obtext.EmbeddedCodeSemNode:
		f, err := os.Open(t.Arg2)
		if err != nil {
			return fmt.Sprintf("Failed to open file: %s", err)
		}
		defer f.Close()
		data, err := io.ReadAll(f)
		if err != nil {
			return fmt.Sprintf("Failed to read file: %s", err)
		}
		return fmt.Sprintf("```%s\n%s\n```\n", t.Arg1, data)
	case *obtext.InlineCodeSemNode:
		return "`" + generateMarkdown(t.Content) + "`"
	case *obtext.UlSemNode:
		out := "\n"
		for _, e := range t.Contents {
			out += " - " + generateMarkdown(e) + "\n"
		}
		return out
	case *obtext.OlSemNode:
		out := "\n"
		for i, e := range t.Contents {
			out += fmt.Sprintf(" %d. %s\n", i+1, generateMarkdown(e))
		}
		return out
	}
	panic(fmt.Sprintf("node type %T was not included in renderer", t))
}
