package obtext

import (
	"fmt"
	"strconv"
)

func Preprocess(node any, allowedObjects []ObjectTmpl) error {
	allowedObjectsMap := make(map[string]ObjectTmpl)
	for _, o := range allowedObjects {
		allowedObjectsMap[o.ObjectType()] = o
	}
	switch node := node.(type) {
	case *Object:
		// Ensure we always do all children first, incase the preprocessing of this relies on the children.
		for _, e := range node.Args {
			if err := Preprocess(e, allowedObjects); err != nil {
				return err
			}
		}
		// Now validate this object
		if tmpl, ok := allowedObjectsMap[node.Type]; !ok {
			return fmt.Errorf("object '%s' was not defined", node.Type)
		} else {
			if err := tmpl.PreprocessArgs(node.Args); err != nil {
				return err
			}
		}
	case *Arg:
		for _, e := range node.Elements {
			if err := Preprocess(e, allowedObjects); err != nil {
				return err
			}
		}
	}
	// Text is always valid
	return nil
}

// Object Template defines how an object should be structured, and contains logic for preprocessing the arguments.
type ObjectTmpl interface {
	ObjectType() string
	PreprocessArgs([]*Arg) error
}

type BasicObjectTmpl struct {
	Type       string
	NumArgs    int
	AllowExtra bool
}

func (t *BasicObjectTmpl) ObjectType() string {
	return t.Type
}

func (t *BasicObjectTmpl) PreprocessArgs(args []*Arg) error {
	if t.AllowExtra {
		if len(args) < t.NumArgs {
			return fmt.Errorf("expected at least %d args, got %d", t.NumArgs, len(args))
		}
	} else {
		if len(args) != t.NumArgs {
			return fmt.Errorf("expected %d args, got %d", t.NumArgs, len(args))
		}
	}
	return nil
}

type CastObjectTmpl struct {
	Type       string
	CastTos    []string
	AllowExtra bool
	CastExtra  string
}

func (t *CastObjectTmpl) ObjectType() string {
	return t.Type
}

func (t *CastObjectTmpl) PreprocessArgs(args []*Arg) error {
	// Check corrent number of args
	if t.AllowExtra {
		if len(args) < len(t.CastTos) {
			return fmt.Errorf("expected at least %d args, got %d", len(t.CastTos), len(args))
		}
	} else {
		if len(args) != len(t.CastTos) {
			return fmt.Errorf("expected %d args, got %d", len(t.CastTos), len(args))
		}
	}
	for i := range args {
		var castTo string
		if i < len(t.CastTos) {
			castTo = t.CastTos[i]
		} else {
			castTo = t.CastExtra
		}
		if castTo == "" {
			continue
		}
		if len(args[i].Elements) != 1 {
			return fmt.Errorf("expected arg %d to be a single element, got %d elements", i, len(args[i].Elements))
		}
		if txt, ok := args[i].Elements[0].(*Text); !ok {
			return fmt.Errorf("expected arg %d to be text, got %T", i, args[i].Elements[0])
		} else {
			switch castTo {
			case "int":
				if casted, err := strconv.Atoi(txt.Value); err != nil {
					return fmt.Errorf("failed to cast arg %d to int: %s", i, err)
				} else {
					args[i].CastValue = casted
				}
			case "float":
				if casted, err := strconv.ParseFloat(txt.Value, 64); err != nil {
					return fmt.Errorf("failed to cast arg %d to float: %s", i, err)
				} else {
					args[i].CastValue = casted
				}
			case "string":
				args[i].CastValue = txt.Value
			default:
				return fmt.Errorf("unknown cast type '%s'", castTo)
			}
		}
	}
	return nil
}
