package obtext

import "github.com/fatih/color"

func Format(o *Object) string {
	return format(o, "", false)
}

func FormatWithAnsiiColors(o *Object) string {
	return format(o, "", true)
}

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
