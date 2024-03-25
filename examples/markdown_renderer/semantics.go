package main

import "github.com/JoshPattman/obtext"

var markDownSemantics = []obtext.ObjectSemantics{
	&obtext.BasicSemantics{Type: "document", NumArgs: 1},
	&obtext.BasicSemantics{Type: "h1", NumArgs: 1},
	&obtext.BasicSemantics{Type: "h2", NumArgs: 1},
	&obtext.BasicSemantics{Type: "ul", AllowExtra: true},
	&obtext.BasicSemantics{Type: "p", NumArgs: 1},
	&obtext.BasicSemantics{Type: "bold", NumArgs: 1},
	&obtext.BasicSemantics{Type: "image", CastArgsTo: []string{"", "string"}},
	&obtext.BasicSemantics{Type: "code", CastArgsTo: []string{"string"}},
}
