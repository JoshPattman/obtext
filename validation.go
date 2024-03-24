package obtext

import "fmt"

type ArgConstraint interface {
	Validate(args []*ObjectArg) error
}

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
	case *ObjectArg:
		for _, e := range node.Elements {
			if err := Validate(e, allowedObjects); err != nil {
				return err
			}
		}
	}
	// Text is always valid
	return nil
}

type NoContraints struct{}

func (NoContraints) Validate(args []*ObjectArg) error {
	return nil
}

type NArgs struct {
	N int
}

func (n NArgs) Validate(args []*ObjectArg) error {
	if len(args) != n.N {
		return fmt.Errorf("expected %d args, got %d", n.N, len(args))
	}
	return nil
}

type AtLeastNArgs struct {
	N int
}

func (n AtLeastNArgs) Validate(args []*ObjectArg) error {
	if len(args) < n.N {
		return fmt.Errorf("expected at least %d args, got %d", n.N, len(args))
	}
	return nil
}

type AtMostNArgs struct {
	N int
}

func (n AtMostNArgs) Validate(args []*ObjectArg) error {
	if len(args) > n.N {
		return fmt.Errorf("expected at most %d args, got %d", n.N, len(args))
	}
	return nil
}

type NoArgs struct{}

func (NoArgs) Validate(args []*ObjectArg) error {
	if len(args) > 0 {
		return fmt.Errorf("expected no args, got %d", len(args))
	}
	return nil
}

type OneArg struct{}

func (OneArg) Validate(args []*ObjectArg) error {
	if len(args) != 1 {
		return fmt.Errorf("expected one arg, got %d", len(args))
	}
	return nil
}
