# Objective Text - A Dead Simple Markup Language
Objective Text is a minimalist markup language designed for blog writing, designed by blending `html` and `latex`. Its key feature is its simplicity, with a grammar that consists of only three elements:

- Objects: Denoted by `@object_name`
- Arguments: Enclosed in `{argument text and objects}`
- Text: Can be any string of characters

Here’s an example of how you might use Objective Text to structure a blog page:

```
@document
{
    @header{ The Advantages of Objective Text }
    @subheader{ Simplicity at its Core }
    @paragraph{ 
        Objective Text stands out for its simplicity. Its grammar comprises just three elements: objects, arguments, and text. This paragraph will have the leading and trailing whitespace removed, so we can put it on a new line for readability.
    }
    @image{A visual representation of Objective Text's grammar}{/path/to/diagram.png}
    @subheader{ A Robust Parser }
    @paragraph{ 
        While it's true that the parser was developed during late-night coding sessions and is my first since a university project two years ago, it's designed to handle nested objects like @italic{this}.
    }
}
```

## Understanding Objects
Objects are the fundamental building blocks of Objective Text. They’re identified by an `@` symbol followed by the object name (e.g. `@header`). Each object is followed by one or more argument blocks, enclosed in `{}`. Whitespace around these blocks is flexible for readability, and leading and trailing whitespace within argument blocks is automatically removed.

One of the unique features of Objective Text is its flexibility: there are no predefined object types. You can define any object, and the parser will process it. How these objects and their arguments are used is determined by the code that interprets the Abstract Syntax Tree (AST). This was actually the inspiration for the language - I wanted to be able to easily add elements such as image galleries to my blog posts, without having to extend a complex parser for markdown.