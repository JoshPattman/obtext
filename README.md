
# Objective Text - A Dead Simple Markup Language

Objective Text is a minimalist markup language designed for blog writing, designed by blending `html` and `latex`. Its key feature is its simplicity, with a grammar that consists of only three elements:

 - Objects: Denoted by @object_name
 - Arguments: Enclosed in {argument text and objects}
 - Text: Can be any string of characters

Here's an example of how you might use Objective Text to structure a blog page:
```
@document {
    @h1 { Section 1: Why ObText Is The Best}
    @p {
        ObText is the best because it is the @bold{best}.
    }
    @img {A picture of ObText being the best} {path/to/img.png}
}
```

## Understanding Objects

Objects are the fundamental building blocks of Objective Text. They're identified by an @ symbol followed by the object name (e.g. @header). Each object is followed by one or more argument blocks, enclosed in {}. Whitespace around these blocks is flexible for readability, and leading and trailing whitespace within argument blocks is automatically removed.

One of the unique features of Objective Text is its flexibility: there are no predefined object types. You can define any object, and the parser will process it. How these objects and their arguments are used is determined by the code that interprets the Abstract Syntax Tree (AST). This was actually the inspiration for the language - I wanted to be able to easily add elements such as image galleries to my blog posts, without having to extend a complex parser for markdown.

## Fun Fact

Did you know, this markdown file was generated from an ObText file! You can see how in ./examples/markdown_renderer
