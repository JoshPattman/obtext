package obtext

import "fmt"

type SemSingleArgNode struct {
	Content *ContentBlockSemNode
}

func (d *SemSingleArgNode) ParseArgs(args []*ContentBlockSemNode) error {
	if len(args) != 1 {
		return fmt.Errorf("document must have exactly 1 child")
	}
	d.Content = args[0]
	return nil
}

type SemCaptionedLinkNode struct {
	CaptionContent *ContentBlockSemNode
	Link           string
}

func (c *SemCaptionedLinkNode) ParseArgs(args []*ContentBlockSemNode) error {
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

type SemListArgNode struct {
	Content []*ContentBlockSemNode
}

func (l *SemListArgNode) ParseArgs(args []*ContentBlockSemNode) error {
	l.Content = args
	return nil
}
