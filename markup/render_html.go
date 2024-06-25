package markup

import (
	"fmt"
	"io"
	"os"

	"github.com/JoshPattman/obtext"
)

// RenderHTML takes a semantic tree using nodes from the markup package and generates an html string from it.
// To customise the html rendering, you can create a copy of this function in your codebase - its only small.
func RenderHTML(t obtext.SemNode, indent string) string {
	switch t := t.(type) {
	case *obtext.ContentBlockSemNode:
		out := ""
		for _, e := range t.Elements {
			out += RenderHTML(e, indent)
		}
		return out
	case *obtext.TextSemNode:
		return indent + t.Text
	case *DocSemNode:
		return indent + RenderHTML(t.Content, indent)
	case *SectionSemNode:
		return "\n" + indent + "<h1>" + RenderHTML(t.Arg1, "") + "</h1>\n" + RenderHTML(t.Arg2, indent+"\t")
	case *SubSectionSemNode:
		return "\n" + indent + "<h2>" + RenderHTML(t.Arg1, "") + "</h2>\n" + RenderHTML(t.Arg2, indent+"\t")
	case *PSemNode:
		return "\n" + indent + "<p>" + RenderHTML(t.Content, "") + "</p>\n"
	case *BoldSemNode:
		return "<b>" + RenderHTML(t.Content, "") + "</b>"
	case *ItalicSemNode:
		return "<i>" + RenderHTML(t.Content, "") + "</i>"
	case *ImageSemNode:
		return "\n" + indent + fmt.Sprintf("<img alt=\"%s\" src=\"%s\" width=50%% align=\"center\"/>\n", RenderHTML(t.CaptionContent, ""), t.Link)
	case *EmbeddedCodeSemNode:
		f, err := os.Open(t.Arg2)
		if err != nil {
			return fmt.Sprintf("Failed to open file: %s", err)
		}
		defer f.Close()
		data, err := io.ReadAll(f)
		if err != nil {
			return fmt.Sprintf("Failed to read file: %s", err)
		}
		return fmt.Sprintf("\n<pre><code>%s</code></pre>\n", data)
	case *InlineCodeSemNode:
		return "<code>" + RenderHTML(t.Content, "") + "</code>"
	case *UlSemNode:
		out := "\n" + indent + "<ul>\n"
		for _, e := range t.Contents {
			out += indent + "\t<li>" + RenderHTML(e, "") + "</li>\n"
		}
		return out + indent + "</ul>"
	case *OlSemNode:
		out := "\n" + indent + "<ol>\n"
		for _, e := range t.Contents {
			out += indent + "\t<li>" + RenderHTML(e, "") + "</li>\n"
		}
		return out + indent + "</ol>"
	}
	panic(fmt.Sprintf("node type %T was not included in renderer", t))
}
