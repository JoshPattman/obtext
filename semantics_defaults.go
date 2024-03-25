package obtext

// DefaultMarkupSemantics is a list of all the default semantic nodes for a markup language using objective text.
// You do not have to use these, and may instead choose to extend it or even create a completely custom set of semantics.
var DefaultMarkupSemantics = []SemNode{
	&DocSemNode{},
	&H1SemNode{},
	&H2SemNode{},
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
	SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (d *DocSemNode) SyntaxType() string {
	return "doc"
}

// H1SemNode is a semantic node that represents a level 1 heading.
type H1SemNode struct {
	SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (h *H1SemNode) SyntaxType() string {
	return "h1"
}

// H2SemNode is a semantic node that represents a level 2 heading.
type H2SemNode struct {
	SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (h *H2SemNode) SyntaxType() string {
	return "h2"
}

// PSemNode is a semantic node that represents a paragraph.
type PSemNode struct {
	SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (p *PSemNode) SyntaxType() string {
	return "p"
}

// BoldSemNode is a semantic node that represents bold text.
type BoldSemNode struct {
	SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (b *BoldSemNode) SyntaxType() string {
	return "b"
}

// ItalicSemNode is a semantic node that represents italic text.
type ItalicSemNode struct {
	SingleArgSemNode
}

// SyntaxType implements the SemNode interface.
func (i *ItalicSemNode) SyntaxType() string {
	return "i"
}

// ImageSemNode is a semantic node that represents an image.
type ImageSemNode struct {
	CaptionedLinkSemNode
}

// SyntaxType implements the SemNode interface.
func (i *ImageSemNode) SyntaxType() string {
	return "img"
}

// VideoSemNode is a semantic node that represents a video.
type VideoSemNode struct {
	CaptionedLinkSemNode
}

// SyntaxType implements the SemNode interface.
func (v *VideoSemNode) SyntaxType() string {
	return "vid"
}

// EmbeddedCodeSemNode is a semantic node that represents embedded code.
type EmbeddedCodeSemNode struct {
	DualStringSemNode
}

// SyntaxType implements the SemNode interface.
func (c *EmbeddedCodeSemNode) SyntaxType() string {
	return "code"
}

// InlineCodeSemNode is a semantic node that represents inline code.
type InlineCodeSemNode struct {
	SingleArgSemNode
}

func (i *InlineCodeSemNode) SyntaxType() string {
	return "icode"
}

// UlSemNode is a semantic node that represents an unordered list.
type UlSemNode struct {
	ListArgSemNode
}

// SyntaxType implements the SemNode interface.
func (u *UlSemNode) SyntaxType() string {
	return "ul"
}

// OlSemNode is a semantic node that represents an ordered list.
type OlSemNode struct {
	ListArgSemNode
}

// SyntaxType implements the SemNode interface.
func (o *OlSemNode) SyntaxType() string {
	return "ol"
}

// LinkSemNode is a semantic node that represents a link.
type LinkSemNode struct {
	CaptionedLinkSemNode
}

// SyntaxType implements the SemNode interface.
func (l *LinkSemNode) SyntaxType() string {
	return "ln"
}
