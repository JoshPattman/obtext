
# Objective Text - A Dead Simple Markup Language

Objective Text is a minimalist markup language designed for blog writing, designed by blending `html` and `latex`. Its key feature is its simplicity, with a grammar that consists of only three elements:

 - Objects: Denoted by `@object_name`
 - Arguments: Enclosed in `{argument text and objects}`
 - Text: Can be any string of characters

Here's an example of how you might use Objective Text to structure a blog page:
```obt
@doc {
    @h1 { Section 1: Why ObText Is The Best}
    @p {
        ObText is the best because it is the @b{best}.
    }
    @img {A picture of ObText being the best} {path/to/img.png}
}
```

For a more detailed example, take a look at `examples/markdown_renderer` example. It shows how you can convert from objective text to markdown, and was actually used to generate this README!

## Understanding Objects

Objects are the fundamental building blocks of Objective Text. They're identified by an @ symbol followed by the object name (e.g. @header). Each object is followed by one or more argument blocks, enclosed in {}. Whitespace around these blocks is flexible for readability, and leading and trailing whitespace within argument blocks is automatically removed.

One of the unique features of Objective Text is its flexibility. During the syntax parsing step, the parser will accept any object with any `@<object type>`, and any number of arguments. However, during the step of converting they syntax tree to a semantic tree, the parser will check if each object is valid given the rules that you give it. Obtext comes with a set of default semantic rules for mark up purposes, `DefaultMarkupSemantics`, but you can easily either modify them, or create your own from scratch. This means that it is trivial to create custom objects, such as an image gallery or a table of contents.

## Usage
```go
package main

import (
	"os"

	"github.com/JoshPattman/obtext"
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
	// Parse the semantics tree according to the rules set in obtext.DefaultMarkupSemantics
	semanticsTree, err := obtext.ParseSem(syntaxTree, obtext.DefaultMarkupSemantics)
	if err != nil {
		panic(err)
	}
	// Do something with the semantics tree
	walk(semanticsTree)
}

func walk(node obtext.SemNode) {
	// ...
}

```
