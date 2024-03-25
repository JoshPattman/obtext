package obtext

// SemNode is an interface that all semantic nodes must implement.
type SemNode interface {
	// SyntaxType returns the type of the node, i.e. '@<node-syntax-type>{...}'.
	// This will be used to match this node with a syntax node.
	SyntaxType() string

	// ParseArgs is called to parse the arguments of the node.
	// It takes a list of ContentBlockSemNode, which are the arguments of the node.
	// For instance, args[1] is the content in the first argument of the node.
	ParseArgs(args []*ContentBlockSemNode) error

	// Children returns all semantic nodes that are children of this node.
	Children() []SemNode
}

// TextSemNode is a special node that is used to represent plain text.
// It should not be passed to ParseSem as one of the semantics.
type TextSemNode struct {
	Text string
}

// SyntaxType implements the SemNode interface.
func (t *TextSemNode) SyntaxType() string {
	return ""
}

// ParseArgs implements the SemNode interface.
func (t *TextSemNode) ParseArgs(parseChild []*ContentBlockSemNode) error {
	panic("cannot parse args for text node")
}

// Children implements the SemNode interface.
func (t *TextSemNode) Children() []SemNode {
	return []SemNode{}
}

// ContentBlockSemNode is a special node that is used to represent a block of content, such as a paragraph or a list,
// basically anything within curly braces in the syntax.
// It should not be passed to ParseSem as one of the semantics.
type ContentBlockSemNode struct {
	Elements []SemNode
}

// ParseArgs implements the SemNode interface.
func (c *ContentBlockSemNode) ParseArgs(args []*ContentBlockSemNode) error {
	panic("cannot parse args for content block node")
}

// SyntaxType implements the SemNode interface.
func (c *ContentBlockSemNode) SyntaxType() string {
	return ""
}

// Children implements the SemNode interface.
func (c *ContentBlockSemNode) Children() []SemNode {
	return c.Elements
}
