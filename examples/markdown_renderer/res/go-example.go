package main

import (
	"os"

	"github.com/JoshPattman/obtext"
	"github.com/JoshPattman/obtext/markup"
)

func main() {
	// First, open the file
	// We will read from an io.Reader in this example, but you can also read from a string or []byte
	f, err := os.Open("path/to/file.obt")
	if err != nil {
		panic(err)
	}
	// Parse the syntax tree
	syntaxTree, err := obtext.ParseSynReader(f)
	if err != nil {
		panic(err)
	}
	// Parse the semantics tree according to the rules set in markup.Semantics
	semanticsTree, err := obtext.ParseSem(syntaxTree, markup.Semantics)
	if err != nil {
		panic(err)
	}
	// Do something with the semantics tree
	walk(semanticsTree)
}

func walk(node obtext.SemNode) {
	// ...
}
