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

func ParseString(data string) (*Object, error) {
	return ParseBytes([]byte(data))
}

func ParseReader(data io.Reader) (*Object, error) {
	buf, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}
	return ParseBytes(buf)
}

func ParseBytes(data []byte) (*Object, error) {
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
	case *Object:
		for _, arg := range n.Args {
			removeWhitespaceOnlyTextFromChildren(arg)
		}
	case *Arg:
		newChildren := make([]Element, 0, len(n.Elements))
		for _, el := range n.Elements {
			if txt, ok := el.(*Text); ok {
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
	case *Object:
		for _, arg := range n.Args {
			stripWhitespaceFromEndChildren(arg)
		}
	case *Arg:
		if len(n.Elements) == 0 {
			return
		}
		if txt, ok := n.Elements[0].(*Text); ok {
			txt.Value = strings.TrimLeft(txt.Value, " \r\n\t")
		}
		if txt, ok := n.Elements[len(n.Elements)-1].(*Text); ok {
			txt.Value = strings.TrimRight(txt.Value, " \r\n\t")
		}
		for _, el := range n.Elements {
			stripWhitespaceFromEndChildren(el)
		}
	}
}

func cleanupEscapedSpecialChars(node any) {
	switch n := node.(type) {
	case *Object:
		for _, arg := range n.Args {
			cleanupEscapedSpecialChars(arg)
		}
	case *Arg:
		for _, el := range n.Elements {
			cleanupEscapedSpecialChars(el)
		}
	case *Text:
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

func tryParseText(data []byte) (*Text, []byte) {
	parsed, groups, remaining := consume(textRegexp, data)
	if !parsed {
		return nil, nil
	}
	return &Text{groups[0]}, remaining
}

func tryParseArg(data []byte) (*Arg, []byte) {
	// Parse some whitespace then a open bracket
	parsed, _, remaining := consume(argRegexpStart, data)
	if !parsed {
		return nil, nil
	}
	data = remaining
	elements := make([]Element, 0)
	for {
		// First try to parse an ending bracket
		parsed, _, remaining = consume(argRegexpEnd, data)
		if parsed {
			// sucsess! return the object arg
			oa := &Arg{Elements: elements}
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

func tryParseObject(data []byte) (*Object, []byte) {
	// Try to parse @obj_type
	parsed, groups, remaining := consume(objRegexpName, data)
	if !parsed {
		return nil, nil
	}
	objType := groups[1]
	data = remaining
	args := make([]*Arg, 0)
	// Parse all remaining args. An arg can be preceded by whitespace (this is dealt with in the arg parser)
	for {
		arg, remaining := tryParseArg(data)
		if arg == nil {
			break
		}
		args = append(args, arg)
		data = remaining
	}
	return &Object{
		Type: objType,
		Args: args,
	}, data
}
