package obtext

// SynElement is either an Object or Text.
type SynElement interface {
	isSynElement()
}

func (ObjectSynNode) isSynElement() {}
func (TextSynNode) isSynElement()   {}

// ObjectSynNode is a syntax node representing an object: @object_name{arg1}{arg2}...
type ObjectSynNode struct {
	Type string
	Args []*ArgSynNode
}

// ArgSynNode is a syntax node representing a list of elements.
// Each element can be either an ObjectSynNode or TextSynNode.
type ArgSynNode struct {
	Elements []SynElement
	// This may be nil, however during validation it is possible that this will get populated according to the validation constraints.
	CastValue any
}

// TextSynNode is a syntax node representing a text value.
type TextSynNode struct {
	Value string
}
