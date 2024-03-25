package obtext

type SemNode interface {
	SyntaxType() string
	ParseArgs(args []*ContentBlockSemNode) error
}

type TextSemNode struct {
	Text string
}

// Text is a special node as it does not match any object, but just plain text
func (t *TextSemNode) SyntaxType() string {
	return ""
}

func (t *TextSemNode) ParseArgs(parseChild []*ContentBlockSemNode) error {
	panic("cannot parse args for text node")
}

// ContentBlockSemNode is a special node that is used to represent a block of content, such as a paragraph or a list.
type ContentBlockSemNode struct {
	Elements []SemNode
}

func (c *ContentBlockSemNode) ParseArgs(args []*ContentBlockSemNode) error {
	panic("cannot parse args for content block node")
}

func (c *ContentBlockSemNode) SyntaxType() string {
	return ""
}
