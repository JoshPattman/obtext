package obtext

// SynElement is either an Object or Text.
type SynElement interface {
	isSynElement()
}

func (ObjectSynNode) isSynElement() {}
func (TextSynNode) isSynElement()   {}

// ObjectSynNode is a named collection of arguments: @object_name{arg1}{arg2}...
type ObjectSynNode struct {
	Type string
	Args []*ArgSynNode
}

// ArgSynNode is a collection of elements, these can be either blocks of text or other objects (or a mix of both).
type ArgSynNode struct {
	Elements []SynElement
	// This may be nil, however during validation it is possible that this will get populated according to the validation constraints.
	CastValue any
}

// TextSynNode is a simple text element that contains just raw text.
type TextSynNode struct {
	Value string
}

func (o *ObjectSynNode) NumArgs() int {
	return len(o.Args)
}

func (a *ArgSynNode) NumElements() int {
	return len(a.Elements)
}
