package markup

import (
	"fmt"
	"io"
	"os"

	"github.com/JoshPattman/obtext"
)

// RenderMarkdown takes a semantic tree using nodes from the markup package and generates a markdown string from it.
// To customise the markdown rendering, you can create a copy of this function in your codebase - its only small.
func RenderMarkdown(t obtext.SemNode) string {
	switch t := t.(type) {
	case *obtext.ContentBlockSemNode:
		out := ""
		for _, e := range t.Elements {
			out += RenderMarkdown(e)
		}
		return out
	case *obtext.TextSemNode:
		return t.Text
	case *DocSemNode:
		return RenderMarkdown(t.Content)
	case *SectionSemNode:
		return "\n# " + RenderMarkdown(t.Arg1) + "\n" + RenderMarkdown(t.Arg2)
	case *SubSectionSemNode:
		return "\n## " + RenderMarkdown(t.Arg1) + "\n" + RenderMarkdown(t.Arg2)
	case *PSemNode:
		return "\n" + RenderMarkdown(t.Content) + "\n"
	case *BoldSemNode:
		return "**" + RenderMarkdown(t.Content) + "**"
	case *ItalicSemNode:
		return "*" + RenderMarkdown(t.Content) + "*"
	case *ImageSemNode:
		return fmt.Sprintf("\n![%s](%s)\n", RenderMarkdown(t.CaptionContent), t.Link)
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
		return fmt.Sprintf("```%s\n%s\n```\n", t.Arg1, data)
	case *InlineCodeSemNode:
		return "`" + RenderMarkdown(t.Content) + "`"
	case *UlSemNode:
		out := "\n"
		for _, e := range t.Contents {
			out += " - " + RenderMarkdown(e) + "\n"
		}
		return out
	case *OlSemNode:
		out := "\n"
		for i, e := range t.Contents {
			out += fmt.Sprintf(" %d. %s\n", i+1, RenderMarkdown(e))
		}
		return out
	}
	panic(fmt.Sprintf("node type %T was not included in renderer", t))
}
