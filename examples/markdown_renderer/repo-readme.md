
<div style="text-align:center"><img alt="icon" src="./icon.png" width=50%/></div>

# Objective Text - A Dead Simple Markup Language

Objective Text (**obtext**) is a minimalist markup language designed for blog writing.

## About

The key features of **obtext** are its extensibility and simplicity, with a grammar that consists of only three elements:
				
 - Objects: Denoted by `@object_name`
 - Arguments: Enclosed in `{argument text and objects}`
 - Text: Can be any string of characters


Objects are the fundamental building blocks of Objective Text. They're identified by an @ symbol followed by the object name (e.g. `@header`). Each object is followed by one or more argument blocks, enclosed in {}. Whitespace around these blocks is flexible for readability, and leading and trailing whitespace within argument blocks is automatically removed.

During the syntax parsing step, the parser will accept any object with any `@<object type>`, and any number of arguments. However, during the step of converting they syntax tree to a semantic tree, the parser will check if each object is valid given the rules that you give it. **obtext** comes with a set of default semantic rules for mark up purposes, `markup.Semantics`, but you can easily either modify them, or create your own from scratch. This means that it is trivial to create custom objects, such as an image gallery or a table of contents.

## Syntax Example

Here's an example of how you might use Objective Text to structure a blog page, using the default rules in the subpackage `markup`:
```obt
@doc {
    @section{ Section 1: Why ObText Is The Best}{
        @para {
            ObText is the best because it is the @bold{best}.
        }
        @img {A picture of ObText being the best} {path/to/img.png}
    }
}
```

For a more detailed example, take a look at both the `examples/markdown_renderer` example, and the renderers provided in `markup`. These tools were actually used to generate this README!

## What Does This Package Contain?

**obtext** is not provided with a single binary or program to perform conversion to other formats (for example **obtext** to markdown converter). It is instead intended to be used within other go programs such as a wesite which renders HTML on the fly. For this reason, **obtext** is provided as a package with parsing functionality that should be called from your specific program.

It is extremely easy to add custom objects to the **obtext** parsing pipeline. To do this, you must first construct a struct that extends the `SemNode` interface. You can either implement all of the methods by hand for this, or you can compose using one of the types defined in `semantics_bases.go` which almost entirely implement a set of common behaviours. You can then pass your struct either in place of, or addition to the other SymNode types present in `markup.Semantics`. Finally, you can add support of your struct to your rendering method. If you are using one of the built-in renderers, you must copy the source code and modify it as you wish. This is done to keep the package as simple and robust as possible, as it is expected that most use-cases of this package will include a custom renderer.
