package obtext

// Element is either an Object or Text.
type Element interface {
	isElement()
}

func (Object) isElement() {}
func (Text) isElement()   {}

// Object is a named collection of arguments: @object_name{arg1}{arg2}...
type Object struct {
	Type string
	Args []*Arg
}

// Arg is a collection of elements, these can be either blocks of text or other objects (or a mix of both).
type Arg struct {
	Elements []Element
	// This may be nil, however during validation it is possible that this will get populated according to the validation constraints.
	CastValue any
}

// Text is a simple text element that contains just raw text.
type Text struct {
	Value string
}

func (o *Object) NumArgs() int {
	return len(o.Args)
}

func (a *Arg) NumElements() int {
	return len(a.Elements)
}
