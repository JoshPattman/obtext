package markup

import "github.com/JoshPattman/obtext"

// DefaultMarkupSemantics is a list of all the default semantic nodes for a markup language using objective text.
// You do not have to use these, and may instead choose to extend it or even create a completely custom set of semantics.
var Semantics = []obtext.SemNode{
	&DocSemNode{},
	&SectionSemNode{},
	&SubSectionSemNode{},
	&PSemNode{},
	&BoldSemNode{},
	&ItalicSemNode{},
	&ImageSemNode{},
	&VideoSemNode{},
	&EmbeddedCodeSemNode{},
	&InlineCodeSemNode{},
	&UlSemNode{},
	&OlSemNode{},
	&LinkSemNode{},
}

// DocSemNode is a semantic node that represents a document.
type DocSemNode struct {
	obtext.SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (d *DocSemNode) SyntaxType() string {
	return "doc"
}

// SectionSemNode is a semantic node that represents a section (level 1 heading usually).
type SectionSemNode struct {
	obtext.DualArgSemNode
}

// SyntaxType implements the SemNode interface.
func (h *SectionSemNode) SyntaxType() string {
	return "section"
}

// SubSectionSemNode is a semantic node that represents a subsection (level 2 heading).
type SubSectionSemNode struct {
	obtext.DualArgSemNode
}

// SyntaxType implements the SemNode interface.
func (h *SubSectionSemNode) SyntaxType() string {
	return "subsection"
}

// PSemNode is a semantic node that represents a paragraph.
type PSemNode struct {
	obtext.SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (p *PSemNode) SyntaxType() string {
	return "para"
}

// BoldSemNode is a semantic node that represents bold text.
type BoldSemNode struct {
	obtext.SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (b *BoldSemNode) SyntaxType() string {
	return "bold"
}

// ItalicSemNode is a semantic node that represents italic text.
type ItalicSemNode struct {
	obtext.SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (i *ItalicSemNode) SyntaxType() string {
	return "italic"
}

// ImageSemNode is a semantic node that represents an image.
type ImageSemNode struct {
	obtext.CaptionedLinkSemNode
}

// SyntaxType implements the SemNode interface.
func (i *ImageSemNode) SyntaxType() string {
	return "img"
}

// VideoSemNode is a semantic node that represents a video.
type VideoSemNode struct {
	obtext.CaptionedLinkSemNode
}

// SyntaxType implements the SemNode interface.
func (v *VideoSemNode) SyntaxType() string {
	return "vid"
}

// EmbeddedCodeSemNode is a semantic node that represents embedded code.
type EmbeddedCodeSemNode struct {
	obtext.DualStringSemNode
}

// SyntaxType implements the SemNode interface.
func (c *EmbeddedCodeSemNode) SyntaxType() string {
	return "code"
}

// InlineCodeSemNode is a semantic node that represents inline code.
type InlineCodeSemNode struct {
	obtext.SingleArgSemNode
}

func (i *InlineCodeSemNode) SyntaxType() string {
	return "icode"
}

// UlSemNode is a semantic node that represents an unordered list.
type UlSemNode struct {
	obtext.ListArgSemNode
}

// SyntaxType implements the SemNode interface.
func (u *UlSemNode) SyntaxType() string {
	return "itemize"
}

// OlSemNode is a semantic node that represents an ordered list.
type OlSemNode struct {
	obtext.ListArgSemNode
}

// SyntaxType implements the SemNode interface.
func (o *OlSemNode) SyntaxType() string {
	return "enumerate"
}

// LinkSemNode is a semantic node that represents a link.
type LinkSemNode struct {
	obtext.CaptionedLinkSemNode
}

// SyntaxType implements the SemNode interface.
func (l *LinkSemNode) SyntaxType() string {
	return "link"
}
