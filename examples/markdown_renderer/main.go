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

	// Try to parse the input file into an AST
	ast, err := obtext.ParseReader(inputFile)
	if err != nil {
		fmt.Println("Failed to parse input file:", err)
		os.Exit(1)
	}

	// Validate the AST (check each object has the correct number of arguments)
	err = obtext.Validate(ast, map[string]obtext.ArgConstraint{
		"document": obtext.OneArg{},       // {content}
		"h1":       obtext.OneArg{},       // {heading text}
		"h2":       obtext.OneArg{},       // {heading text}
		"ul":       obtext.NoContraints{}, // ul can have as many args as in the list
		"p":        obtext.OneArg{},       // {text}
		"bold":     obtext.OneArg{},       // {text}
		"image":    obtext.NArgs{N: 2},    // {alt-text}{url}
		"code":     obtext.OneArg{},
	})
	if err != nil {
		fmt.Println("Failed to validate input file:", err)
		os.Exit(1)
	}

	// Pretty print the AST to stdout if requested
	if prettyPrint {
		fmt.Println(obtext.FormatWithAnsiiColors(ast))
	}

	// Create the output file
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Generate the markdown from the AST and write it to the output file
	md := generateMarkdown(ast)
	fmt.Fprint(outputFile, md)

	fmt.Println("Successfully wrote markdown to", outputFileName)
}

// generateMarkdown generates markdown from the given AST
func generateMarkdown(t any) string {
	switch t := t.(type) {
	case *obtext.Object:
		// If this node is an object, choose what to do with it based on its type
		switch t.Type {
		case "document":
			return generateMarkdown(t.Args[0])
		case "h1":
			return "\n# " + generateMarkdown(t.Args[0]) + "\n"
		case "h2":
			return "\n## " + generateMarkdown(t.Args[0]) + "\n"
		case "p":
			return "\n" + generateMarkdown(t.Args[0]) + "\n"
		case "bold":
			return "**" + generateMarkdown(t.Args[0]) + "**"
		case "image":
			return fmt.Sprintf("\n![%s](%s)\n", generateMarkdown(t.Args[0]), generateMarkdown(t.Args[1]))
		case "ul":
			out := "\n"
			for _, e := range t.Args {
				out += " - " + generateMarkdown(e) + "\n"
			}
			return out
		case "code":
			f, err := os.Open(generateMarkdown(t.Args[0]))
			if err != nil {
				return fmt.Sprintf("Failed to open file: %s", err)
			}
			defer f.Close()
			data, err := io.ReadAll(f)
			if err != nil {
				return fmt.Sprintf("Failed to read file: %s", err)
			}
			return fmt.Sprintf("```\n%s\n```\n", data)
		default:
			panic("unknown object type")
		}
	case *obtext.Arg:
		// If this node is an object argument, generate markdown for each element and append it together
		out := ""
		for _, e := range t.Elements {
			out += generateMarkdown(e)
		}
		return out
	case *obtext.Text:
		// If this node is text, simply return the text value
		return t.Value
	default:
		panic("unknown type")
	}
}
