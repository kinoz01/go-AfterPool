# Introduction

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A string literal is defined between double quotes:

```go
const name = "Jane"
```

Strings can be concatenated via the `+` operator:

```go
"Jane" + " " + "Austen"
// => "Jane Austen"
```

Some special characters need to be escaped with a leading backslash, such as `\t` for a tab and `\n` for a new line in strings.

```go
"How is the weather today?\nIt's sunny"  
// =>
// How is the weather today?
// It's sunny
```

The `strings` package contains many useful functions to work on strings.
For more information about string functions, check out the [strings package documentation](https://pkg.go.dev/strings).
Here are some examples:

```go
import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
strings.ToLower("MaKEmeLoweRCase")
// => "makemelowercase"

// strings.Repeat returns a string with a substring given as argument repeated many times
strings.Repeat("Go", 3)
// => "GoGoGo"
```

## Special characters

As mentioned previously, some special characters need to be escaped with a leading backslash.
Below is a more detailed list of what those special characters are:

| Value | Description          |
| ----- | -------------------- |
| `\a`  | Alert or bell        |
| `\b`  | Backspace            |
| `\\`  | Backslash            |
| `\t`  | Horizontal tab       |
| `\n`  | Line feed or newline |
| `\f`  | Form feed            |
| `\r`  | Carriage return      |
| `\v`  | Vertical tab         |
| `\'`  | Single quote         |
| `\"`  | Double quote         |

```go
fmt.Println("Joe\nWilliam\nJack\nAverell") 

// Output:
// Joe
// William
// Jack
// Averell
```

## Raw string literals

Raw string literals use backticks (\`) as their delimiter instead of double quotes and all characters in it are interpreted literally, meaning that there is no need to escape characters or newlines.

This is an example of a multiline string:

```go
const daltons = `Joe
William
Jack
Averell`
```

This string has multiple lines. More specifically, it has 3 `\n`. However, because we are a raw string literal, we don't need to put the `\n` explicitly in the string. The newlines will be interpreted literally.

Here are some other examples comparing raw string literals with regular string literals:

```go
`abc`
// same as "abc"

"\"" // regular string literal with 1 character: a double quote
// same as `"` a raw string literal with 1 character: a double quote

`\n
` // raw string literal with 3 character: a backslash, an 'n', and a newline
// same as "\\n\n"  a regular string literal with 3 characters: a backslash, an 'n', and a newline

"\t\n" // regular string literal with 2 characters: a tab and a newline
`\t\n`// raw string literal with 4 characters: two backslashes, a 't', and an 'n'
```

# Exercice Instructions

There is an appliance store called "Tech Palace" nearby.
The owner of the store recently installed a big display to use for marketing messages and to show a special greeting when customers scan their loyalty cards at the entrance.
The display consists of lots of small LED lights and can show multiple lines of text.

The store owner needs your help with the code that is used to generate the text for the new display.

## 1. Create the welcome message

For most customers who scan their loyalty cards, the store owner wants to see `Welcome to the Tech Palace, ` followed by the name of the customer in capital letters on the display.

Implement the function `WelcomeMessage` that accepts the name of the customer as a `string` argument and returns the desired message as a `string`.

```go
WelcomeMessage("Judy")
// => Welcome to the Tech Palace, JUDY
```

## 2. Add a fancy border

For loyal customers that buy a lot at the store, the owner wants the welcome display to be more fancy by adding a line of stars before and after the welcome message.
They are not sure yet how many stars should be in the lines so they want that to be configurable.

Write a function `AddBorder` that accepts a welcome message (a `string`) and the number of stars per line (type `int`) as arguments.
It should return a `string` that consists of 3 lines, a line with the desired number of stars, then the welcome message as it was passed in, then another line of stars.

```go
AddBorder("Welcome!", 10)
```

Should return the following:

```
**********
Welcome!
**********
```

## 3. Clean up old marketing messages

Before installing this new display, the store had a similar display that could only show non-configurable, static lines.
The owner would like to reuse some of the old marketing messages on the new display.
However, the data already includes a star border and some unfortunate whitespaces.
Your task is to clean up the messages so they can be re-used.

Implement a function `CleanUpMessage` that accepts the old marketing message as a string.
The function should first remove all stars from the text and afterwards remove the leading and trailing whitespaces from the remaining text.
The function should then return the cleaned up message.

```go
message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

CleanUpMessage(message)
// => BUY NOW, SAVE 10%
```