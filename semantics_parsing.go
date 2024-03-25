package obtext

import (
	"fmt"
	"reflect"
)

// ParseSem parses the given syntax tree and returns the semantics tree, or an error if the syntax tree is invalid.
// It parses based on the given semantics, which is a list of all possible semantic nodes.
// Each node contains information about what @<syntax-type> it should match, and how to parse its arguments.
func ParseSem(node any, semantics []SemNode) (SemNode, error) {
	semanticLookup := make(map[string]SemNode)
	for _, o := range semantics {
		semanticLookup[o.SyntaxType()] = o
	}
	return parseSem(node, semanticLookup)
}

func parseSem(node any, semanticLookup map[string]SemNode) (SemNode, error) {
	switch node := node.(type) {
	case *ObjectSynNode:
		if sem, ok := semanticLookup[node.Type]; !ok {
			return nil, fmt.Errorf("object '%s' was not defined", node.Type)
		} else {
			// First parse all children of all args
			parsedArgs := make([]*ContentBlockSemNode, len(node.Args))
			for i, arg := range node.Args {
				parsedArgs[i] = &ContentBlockSemNode{Elements: make([]SemNode, len(arg.Elements))}
				for j, e := range arg.Elements {
					parsed, err := parseSem(e, semanticLookup)
					if err != nil {
						return nil, err
					}
					parsedArgs[i].Elements[j] = parsed
				}
			}
			// Now using our new semantic children args, parse the object
			newNode := reflect.New(reflect.TypeOf(sem).Elem()).Interface().(SemNode)
			if err := newNode.ParseArgs(parsedArgs); err != nil {
				return nil, err
			}
			return newNode, nil
		}
	case *TextSynNode:
		return &TextSemNode{Text: node.Value}, nil

	}
	panic("unknown type")
}
