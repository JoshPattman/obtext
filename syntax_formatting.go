package obtext

import "github.com/fatih/color"

// FormatSyn returns a string representation of the syntax tree object with nice indentation.
func FormatSyn(o *ObjectSynNode) string {
	return formatSyn(o, "", false)
}

// FormatSynWithAnsiiColors returns a string representation of the syntax tree object with nice indentation and ansii colors.
func FormatSynWithAnsiiColors(o *ObjectSynNode) string {
	return formatSyn(o, "", true)
}

// formatSyn is a recursive function that formats the object with indentation and optionally colors.
func formatSyn(o *ObjectSynNode, indent string, withAnsiiColors bool) string {
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
			case *ObjectSynNode:
				out += formatSyn(e, indent+"  ", withAnsiiColors)
			case *TextSynNode:
				out += indent + "  " + e.Value + "\n"
			}
		}
		out += indent + conditionalColString("}\n", color.YellowString)
	}
	return out
}
