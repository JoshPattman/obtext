package main

import (
	"fmt"
	"os"
	"time"

	"github.com/JoshPattman/obtext"
)

func main() {
	f, err := os.Open("test.obt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	startParseTime := time.Now()
	ast, ok := obtext.ParseReader(f)
	if !ok {
		panic("failed to parse file")
	}

	fmt.Printf("Read Obtext file in: %v:\n", time.Since(startParseTime))

	err = obtext.Validate(ast, map[string]obtext.ArgConstraint{
		"document": obtext.OneArg{},
		"h1":       obtext.OneArg{},
		"h2":       obtext.OneArg{},
		"p":        obtext.OneArg{},
		"bold":     obtext.OneArg{},
		"image":    obtext.NArgs{N: 2}, // Image has {alt-text}{url}
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(obtext.FormatWithAnsiiColors(ast))

	md := generateMarkdown(ast)
	f3, err := os.Create("test.md")
	if err != nil {
		panic(err)
	}
	defer f3.Close()
	fmt.Fprint(f3, md)
}

func generateMarkdown(t any) string {
	switch t := t.(type) {
	case *obtext.Object:
		switch t.Name {
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
		default:
			panic("unknown object type")
		}
	case *obtext.ObjectArg:
		out := ""
		for _, e := range t.Elements {
			out += generateMarkdown(e)
		}
		return out
	case *obtext.Text:
		return t.Value
	default:
		panic("unknown type")
	}
}
