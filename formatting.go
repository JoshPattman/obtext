package obtext

import "github.com/fatih/color"

// Format returns a string representation of the object with nice indentation.
func Format(o *Object) string {
	return format(o, "", false)
}

// FormatWithAnsiiColors returns a string representation of the object with nice indentation and ansii colors.
func FormatWithAnsiiColors(o *Object) string {
	return format(o, "", true)
}

// format is a recursive function that formats the object with indentation and optionally colors.
func format(o *Object, indent string, withAnsiiColors bool) string {
	conditionalColString := func(s string, f func(string, ...any) string) string {
		if withAnsiiColors {
			return f(s)
		}
		return s
	}
	out := indent + conditionalColString("@"+o.Type, color.BlueString) + "\n"
	for _, a := range o.Args {
		out += indent + conditionalColString("{\n", color.YellowString)
		for _, e := range a.Elements {
			switch e := e.(type) {
			case *Object:
				out += format(e, indent+"  ", withAnsiiColors)
			case *Text:
				out += indent + "  " + e.Value + "\n"
			}
		}
		out += indent + conditionalColString("}\n", color.YellowString)
	}
	return out
}
