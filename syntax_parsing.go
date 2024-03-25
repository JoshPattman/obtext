package obtext

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

var textRegexp = regexp.MustCompile(`(?:\\[@}]|[^@}])+`)

var argRegexpStart = regexp.MustCompile("[ \n\r\t]*{")
var argRegexpEnd = regexp.MustCompile("}")

var objRegexpName = regexp.MustCompile("@([a-zA-Z0-9_]+)")

// ParseSynString is a convenience function that calls ParseBytes after converting the string to a byte slice
func ParseSynString(data string) (*ObjectSynNode, error) {
	return ParseSynBytes([]byte(data))
}

// ParseSynReader is a convenience function that reads all data from the reader and then calls ParseBytes
func ParseSynReader(data io.Reader) (*ObjectSynNode, error) {
	buf, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}
	return ParseSynBytes(buf)
}

// ParseSynBytes parses the given byte slice and returns the AST, or an error if the data is invalid.
// The resulting AST represents only they syntax, and should probably not be used directly.
// Instead, you should call ParseSem on the result to parse the syntax tree into a semantics tree.
func ParseSynBytes(data []byte) (*ObjectSynNode, error) {
	// Initially, trim all whitespace from front and end (some editors add a newline at the end)
	data = []byte(strings.Trim(string(data), " \r\n\t"))
	// First, try to parse the messy ast
	obj, remaining := tryParseObject(data)
	if obj == nil {
		return nil, fmt.Errorf("failed to parse: invalid syntax")
	}
	if len(remaining) > 0 {
		return nil, fmt.Errorf("failed to parse: remaining characters detected")
	}
	// Now traverse the tree, removing all text that is only whitespace
	removeWhitespaceOnlyTextFromChildren(obj)
	// Now traverse the tree, and:
	// - trim whitespace from the front of any text elements that are the first child of an object arg
	// - trim whitespace from the back of any text elements that are the last child of an object arg
	stripWhitespaceFromEndChildren(obj)
	// Finally, remove the escaping around any escaped special characters
	cleanupEscapedSpecialChars(obj)
	return obj, nil
}

func removeWhitespaceOnlyTextFromChildren(node any) {
	switch n := node.(type) {
	case *ObjectSynNode:
		for _, arg := range n.Args {
			removeWhitespaceOnlyTextFromChildren(arg)
		}
	case *ArgSynNode:
		newChildren := make([]SynElement, 0, len(n.Elements))
		for _, el := range n.Elements {
			if txt, ok := el.(*TextSynNode); ok {
				if strings.Trim(txt.Value, " \r\n\t") == "" {
					continue
				}
			}
			newChildren = append(newChildren, el)
		}
		n.Elements = newChildren
		for _, el := range n.Elements {
			removeWhitespaceOnlyTextFromChildren(el)
		}
	}
}

func stripWhitespaceFromEndChildren(node any) {
	switch n := node.(type) {
	case *ObjectSynNode:
		for _, arg := range n.Args {
			stripWhitespaceFromEndChildren(arg)
		}
	case *ArgSynNode:
		if len(n.Elements) == 0 {
			return
		}
		if txt, ok := n.Elements[0].(*TextSynNode); ok {
			txt.Value = strings.TrimLeft(txt.Value, " \r\n\t")
		}
		if txt, ok := n.Elements[len(n.Elements)-1].(*TextSynNode); ok {
			txt.Value = strings.TrimRight(txt.Value, " \r\n\t")
		}
		for _, el := range n.Elements {
			stripWhitespaceFromEndChildren(el)
		}
	}
}

func cleanupEscapedSpecialChars(node any) {
	switch n := node.(type) {
	case *ObjectSynNode:
		for _, arg := range n.Args {
			cleanupEscapedSpecialChars(arg)
		}
	case *ArgSynNode:
		for _, el := range n.Elements {
			cleanupEscapedSpecialChars(el)
		}
	case *TextSynNode:
		n.Value = strings.ReplaceAll(n.Value, "\\@", "@")
		n.Value = strings.ReplaceAll(n.Value, "\\}", "}")
		n.Value = strings.ReplaceAll(n.Value, "\\{", "{")
	}
}

func consume(reg *regexp.Regexp, data []byte) (bool, []string, []byte) {
	// Try to match the regular expression (only get the first match)
	locs := reg.FindSubmatchIndex(data)
	if locs == nil {
		return false, nil, nil
	}
	if locs[0] != 0 {
		return false, nil, nil
	}

	groups := make([]string, len(locs)/2)
	for i := 0; i < len(locs)/2; i++ {
		groups[i] = string(data[locs[2*i]:locs[2*i+1]])
	}
	return true, groups, data[locs[1]:]
}

func tryParseText(data []byte) (*TextSynNode, []byte) {
	parsed, groups, remaining := consume(textRegexp, data)
	if !parsed {
		return nil, nil
	}
	return &TextSynNode{groups[0]}, remaining
}

func tryParseArg(data []byte) (*ArgSynNode, []byte) {
	// Parse some whitespace then a open bracket
	parsed, _, remaining := consume(argRegexpStart, data)
	if !parsed {
		return nil, nil
	}
	data = remaining
	elements := make([]SynElement, 0)
	for {
		// First try to parse an ending bracket
		parsed, _, remaining = consume(argRegexpEnd, data)
		if parsed {
			// sucsess! return the object arg
			oa := &ArgSynNode{Elements: elements}
			return oa, remaining
		}
		// Now try to parse a new object
		obj, remaining := tryParseObject(data)
		if obj != nil {
			data = remaining
			elements = append(elements, obj)
			continue
		}

		// Finally try to parse some text
		txt, remaining := tryParseText(data)
		if txt != nil {
			data = remaining
			elements = append(elements, txt)
			continue
		}

		// If all of those failed, this is not parseable, so return nil
		return nil, nil
	}
}

func tryParseObject(data []byte) (*ObjectSynNode, []byte) {
	// Try to parse @obj_type
	parsed, groups, remaining := consume(objRegexpName, data)
	if !parsed {
		return nil, nil
	}
	objType := groups[1]
	data = remaining
	args := make([]*ArgSynNode, 0)
	// Parse all remaining args. An arg can be preceded by whitespace (this is dealt with in the arg parser)
	for {
		arg, remaining := tryParseArg(data)
		if arg == nil {
			break
		}
		args = append(args, arg)
		data = remaining
	}
	return &ObjectSynNode{
		Type: objType,
		Args: args,
	}, data
}
