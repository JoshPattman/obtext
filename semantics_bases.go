package obtext

import (
	"fmt"
)

// SingleArgSemNode is a semantic node that has exactly 1 child, which is a content block.
// It only partially implements the SemNode interface, as it does not implement SyntaxType.
type SingleArgSemNode struct {
	Content *ContentBlockSemNode
}

// ParseArgs implements the SemNode interface.
func (d *SingleArgSemNode) ParseArgs(args []*ContentBlockSemNode) error {
	if len(args) != 1 {
		return fmt.Errorf("document must have exactly 1 child")
	}
	d.Content = args[0]
	return nil
}

// Children implements the SemNode interface.
func (d *SingleArgSemNode) Children() []SemNode {
	return []SemNode{d.Content}
}

// CaptionedSemNode is a semantic node that has exactly 2 children, the first of which is a content block representing a caption,
// and the second which is a string representing a URL.
// It only partially implements the SemNode interface, as it does not implement SyntaxType.
type CaptionedLinkSemNode struct {
	CaptionContent *ContentBlockSemNode
	Link           string
}

// ParseArgs implements the SemNode interface.
func (c *CaptionedLinkSemNode) ParseArgs(args []*ContentBlockSemNode) error {
	if len(args) != 2 {
		return fmt.Errorf("link must have exactly 2 children")
	}
	if len(args[1].Elements) != 1 {
		return fmt.Errorf("link arg1 must have exactly 1 child")
	}
	text, ok := args[1].Elements[0].(*TextSemNode)
	if !ok {
		return fmt.Errorf("link arg1 must be text")
	}
	c.Link = text.Text
	c.CaptionContent = args[0]
	return nil
}

// Children implements the SemNode interface.
func (c *CaptionedLinkSemNode) Children() []SemNode {
	return []SemNode{c.CaptionContent}
}

// DualStringSemNode is a semantic node that has exactly 2 children, both of which are strings.
// It only partially implements the SemNode interface, as it does not implement SyntaxType.
type DualStringSemNode struct {
	Arg1 string
	Arg2 string
}

// ParseArgs implements the SemNode interface.
func (d *DualStringSemNode) ParseArgs(args []*ContentBlockSemNode) error {
	if len(args) != 2 {
		return fmt.Errorf("dual string must have exactly 2 children")
	}
	if len(args[0].Elements) != 1 {
		return fmt.Errorf("dual string arg1 must have exactly 1 child")
	}
	if len(args[1].Elements) != 1 {
		return fmt.Errorf("dual string arg2 must have exactly 1 child")
	}
	arg1, ok := args[0].Elements[0].(*TextSemNode)
	if !ok {
		return fmt.Errorf("dual string arg1 must be text")
	}
	arg2, ok := args[1].Elements[0].(*TextSemNode)
	if !ok {
		return fmt.Errorf("dual string arg2 must be text")
	}
	d.Arg1 = arg1.Text
	d.Arg2 = arg2.Text
	return nil
}

// Children implements the SemNode interface.
func (d *DualStringSemNode) Children() []SemNode {
	return []SemNode{}
}

// ListArgSemNode is a semantic node that has a list of children, each of which is a content block.
// It only partially implements the SemNode interface, as it does not implement SyntaxType.
type ListArgSemNode struct {
	Contents []*ContentBlockSemNode
}

// ParseArgs implements the SemNode interface.
func (l *ListArgSemNode) ParseArgs(args []*ContentBlockSemNode) error {
	l.Contents = args
	return nil
}

// Children implements the SemNode interface.
func (l *ListArgSemNode) Children() []SemNode {
	res := make([]SemNode, len(l.Contents))
	for i, c := range l.Contents {
		res[i] = c
	}
	return res
}

// DualArgSemNode is a semantic node that has exactly 2 children, both of which are futher content.
type DualArgSemNode struct {
	Arg1 *ContentBlockSemNode
	Arg2 *ContentBlockSemNode
}

// ParseArgs implements the SemNode interface.
func (d *DualArgSemNode) ParseArgs(args []*ContentBlockSemNode) error {
	if len(args) != 2 {
		return fmt.Errorf("dual arg must have exactly 2 children")
	}
	d.Arg1 = args[0]
	d.Arg2 = args[1]
	return nil
}

// Children implements the SemNode interface.
func (d *DualArgSemNode) Children() []SemNode {
	return []SemNode{d.Arg1, d.Arg2}
}
