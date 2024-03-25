package obtext

var DefaultMarkupSemantics = []SemNode{
	&DocSemNode{},
	&H1SemNode{},
	&H2SemNode{},
	&PSemNode{},
	&BoldSemNode{},
	&ItalicSemNode{},
	&ImageSemNode{},
	&EmbeddedCodeSemNode{},
	&UlSemNode{},
	&OlSemNode{},
}

// Document type
type DocSemNode struct {
	SemSingleArgNode
}

func (d *DocSemNode) SyntaxType() string {
	return "doc"
}

// H1 type
type H1SemNode struct {
	SemSingleArgNode
}

func (h *H1SemNode) SyntaxType() string {
	return "h1"
}

// H2 type
type H2SemNode struct {
	SemSingleArgNode
}

func (h *H2SemNode) SyntaxType() string {
	return "h2"
}

// Paragraph type
type PSemNode struct {
	SemSingleArgNode
}

func (p *PSemNode) SyntaxType() string {
	return "p"
}

// Bold type
type BoldSemNode struct {
	SemSingleArgNode
}

func (b *BoldSemNode) SyntaxType() string {
	return "b"
}

// Italic type
type ItalicSemNode struct {
	SemSingleArgNode
}

func (i *ItalicSemNode) SyntaxType() string {
	return "i"
}

// Image type
type ImageSemNode struct {
	SemCaptionedLinkNode
}

func (i *ImageSemNode) SyntaxType() string {
	return "img"
}

// Code type
type EmbeddedCodeSemNode struct {
	SemCaptionedLinkNode
}

func (c *EmbeddedCodeSemNode) SyntaxType() string {
	return "code"
}

// Ul type
type UlSemNode struct {
	SemListArgNode
}

func (u *UlSemNode) SyntaxType() string {
	return "ul"
}

// Ol type
type OlSemNode struct {
	SemListArgNode
}

func (o *OlSemNode) SyntaxType() string {
	return "ol"
}
