package obtext

import (
	"fmt"
	"testing"
	"time"
)

func TestParsingByManualInspection(t *testing.T) {
	startParseTime := time.Now()
	ast, err := ParseString("@doc{@h1{First header}\n\n@h2{ \nSecond Header\n }\n@p{this is a paragraph with inline @b{bold} inside of it.}}")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Parsed in ", time.Since(startParseTime))
	fmt.Println(FormatWithAnsiiColors(ast))
}
