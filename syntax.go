package obtext

type Element interface {
	isElement()
}

func (Object) isElement() {}
func (Text) isElement()   {}

// Object is a named collection of arguments: @object_name{arg1}{arg2}...
type Object struct {
	Type string
	Args []*ObjectArg
}

// ObjectArg is a collection of elements, these can be either blocks of text or other objects (or a mix of both).
type ObjectArg struct {
	Elements []Element
}

// Text is a simple text element that contains just raw text.
type Text struct {
	Value string
}

func (o *Object) NumArgs() int {
	return len(o.Args)
}

func (oa *ObjectArg) NumElements() int {
	return len(oa.Elements)
}
