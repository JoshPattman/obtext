package obtext

import "fmt"

// ArgConstraint is an interface that can be implemented to validate the arguments of an object.
type ArgConstraint interface {
	// Validate checks if the given arguments are valid according to this constraint, and returns an error explaining why if not.
	Validate(args []*Arg) error
}

// Validate looks through the ast and checks:
//
//   - All objects in the ast are defined in allowedObjects
//   - All objects have valid arguments according to the constraints in allowedObjects
//
// You can also extend the behaviour by adding custom constraints, as ArgConstraint is an interface.
func Validate(node any, allowedObjects map[string]ArgConstraint) error {
	switch node := node.(type) {
	case *Object:
		if validator, ok := allowedObjects[node.Type]; !ok {
			return fmt.Errorf("object '@%s' was not defined, did you add it to allowed objects?", node.Type)
		} else {
			if err := validator.Validate(node.Args); err != nil {
				return err
			}
		}
		for _, e := range node.Args {
			if err := Validate(e, allowedObjects); err != nil {
				return err
			}
		}
	case *Arg:
		for _, e := range node.Elements {
			if err := Validate(e, allowedObjects); err != nil {
				return err
			}
		}
	}
	// Text is always valid
	return nil
}

// NoContraints is a constraint that allows any number of arguments with any content in them.
type NoContraints struct{}

func (NoContraints) Validate(args []*Arg) error {
	return nil
}

// NoArgs is a constraint that requires no arguments.
type NoArgs struct{}

func (NoArgs) Validate(args []*Arg) error {
	if len(args) > 0 {
		return fmt.Errorf("expected no args, got %d", len(args))
	}
	return nil
}

// OneArg is a constraint that requires exactly one argument.
type OneArg struct{}

func (OneArg) Validate(args []*Arg) error {
	if len(args) != 1 {
		return fmt.Errorf("expected one arg, got %d", len(args))
	}
	return nil
}

// NArgs is a constraint that requires exactly N arguments.
type NArgs struct {
	N int
}

func (n NArgs) Validate(args []*Arg) error {
	if len(args) != n.N {
		return fmt.Errorf("expected %d args, got %d", n.N, len(args))
	}
	return nil
}

// AtLeastNArgs is a constraint that requires at least N arguments.
type AtLeastNArgs struct {
	N int
}

func (n AtLeastNArgs) Validate(args []*Arg) error {
	if len(args) < n.N {
		return fmt.Errorf("expected at least %d args, got %d", n.N, len(args))
	}
	return nil
}

// AtMostNArgs is a constraint that requires at most N arguments.
type AtMostNArgs struct {
	N int
}

func (n AtMostNArgs) Validate(args []*Arg) error {
	if len(args) > n.N {
		return fmt.Errorf("expected at most %d args, got %d", n.N, len(args))
	}
	return nil
}
