package obtext

import (
	"fmt"
	"strconv"
)

// ProcessSemantics takes an AST and a list of ObjectTmpls, and checks that the AST is semantically correct, according to the given object templates.
// Any objects that are not given, or who fail the template check given by their template, will return an error.
func ProcessSemantics(node any, templates []ObjectTmpl) error {
	allowedObjects := make(map[string]ObjectTmpl)
	for _, o := range templates {
		allowedObjects[o.ObjectType()] = o
	}
	switch node := node.(type) {
	case *Object:
		// Ensure we always do all children first, incase the preprocessing of this relies on the children.
		for _, e := range node.Args {
			if err := ProcessSemantics(e, templates); err != nil {
				return err
			}
		}
		// Now validate this object
		if tmpl, ok := allowedObjects[node.Type]; !ok {
			return fmt.Errorf("object '%s' was not defined", node.Type)
		} else {
			if err := tmpl.ProcessArgSemantics(node.Args); err != nil {
				return err
			}
		}
	case *Arg:
		for _, e := range node.Elements {
			if err := ProcessSemantics(e, templates); err != nil {
				return err
			}
		}
	}
	// Text is always valid
	return nil
}

// ObjectTmpl is an iterface, which serves two purposes:
//
//   - It tells the preprocessor that the type of object @Type is allowed
//   - It provides the preprocessor with a function to process the arguments of the object (for example, check there are the correct number, check that arg1 is just a text block that can be cast to an int)
type ObjectTmpl interface {
	ObjectType() string
	ProcessArgSemantics([]*Arg) error
}

// BasicTmpl is a simple implementation of ObjectTmpl, which allows you to specify the number of arguments, the types of the arguments, and how to cast them.
// It also allows extra args, but these can only be cast to one type.
type BasicTmpl struct {
	Type        string
	NumArgs     int
	CastArgsTo  []string
	AllowExtra  bool
	CastExtraTo string
	Caster      Caster
}

func (t *BasicTmpl) ObjectType() string {
	return t.Type
}

func (t *BasicTmpl) ProcessArgSemantics(args []*Arg) error {
	// If we have a set of casting rules, then override the number of args (means you don't have to specify it)
	if t.CastArgsTo != nil {
		t.NumArgs = len(t.CastArgsTo)
	}
	// If we did not specify a caster, use the default caster
	if t.Caster == nil {
		t.Caster = &DefaultCaster{}
	}
	// Check corrent number of args
	if t.AllowExtra {
		if len(args) < t.NumArgs {
			return fmt.Errorf("expected at least %d args for object '%s', got %d", t.NumArgs, t.ObjectType(), len(args))
		}
	} else {
		if len(args) != t.NumArgs {
			return fmt.Errorf("expected %d args for object '%s', got %d", t.NumArgs, t.ObjectType(), len(args))
		}
	}
	for i := range args {
		var castTo string
		if i < t.NumArgs {
			// Skip if we did not specify a set of casting rules
			if t.CastArgsTo == nil {
				continue
			}
			castTo = t.CastArgsTo[i]
		} else {
			castTo = t.CastExtraTo
		}
		// Skip if we did not specify a cast type (dont cast)
		if castTo == "" {
			continue
		}
		if len(args[i].Elements) != 1 {
			return fmt.Errorf("expected arg %d of object '%s' to be a single element, got %d elements", i, t.ObjectType(), len(args[i].Elements))
		}
		if txt, ok := args[i].Elements[0].(*Text); !ok {
			return fmt.Errorf("expected arg %d to be text, got %T", i, args[i].Elements[0])
		} else {
			v, err := t.Caster.CastTo(txt.Value, castTo)
			if err != nil {
				return err
			}
			args[i].CastValue = v
		}
	}
	return nil
}

type Caster interface {
	CastTo(s string, to string) (any, error)
}

type DefaultCaster struct{}

func (d *DefaultCaster) CastTo(s string, to string) (any, error) {
	switch to {
	case "int":
		return strconv.Atoi(s)
	case "float":
		return strconv.ParseFloat(s, 64)
	case "string":
		return s, nil
	default:
		return nil, fmt.Errorf("unknown cast type '%s'", to)
	}
}
