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

func (a *Arg) IsPureText() bool {
	if len(a.Elements) == 1 {
		_, ok := a.Elements[0].(*Text)
		return ok
	}
	return false
}

func (a *Arg) PureText() string {
	if a.IsPureText() {
		return a.Elements[0].(*Text).Value
	}
	panic("arg is not just text, please check before calling JustText() by using IsJustText()")
}
