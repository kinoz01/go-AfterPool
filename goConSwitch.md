# Introduction

Like other languages, Go also provides a `switch` statement.
Switch statements are a shorter way to write long `if ... else if` statements.
To make a switch, we start by using the keyword `switch` followed by a value or expression.
We then declare each one of the conditions with the `case` keyword.
We can also declare a `default` case, that will run when none of the previous `case` conditions matched:

```go
operatingSystem := "windows"

switch operatingSystem {
case "windows":
    // do something if the operating system is windows
case "linux":
    // do something if the operating system is linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
}
```
## Boolean Switch

One interesting thing about switch statements, is that the value after the `switch` keyword can be omitted, and we can have boolean conditions for each `case`:

```go
age := 21

switch {
case age > 20 && age < 30:
    // do something if age is between 20 and 30
case age == 10:
    // do something if age is equal to 10
default:
    // do something else for every other case
}
```

For cases where you'd like to use || (OR logic), you can list multiple values in a single case separated by commas. This is equivalent to checking if the switch expression matches any one of the listed values. Here’s how you can do it:

```go
    value := "apple"
    switch value {
    case "apple", "banana":
        fmt.Println("Fruit is either an apple or a banana")
    case "carrot", "beet":
        fmt.Println("Vegetable is either a carrot or a beet.")
    default:
        fmt.Println("Unknown category.")
    }
```

## Nested Switch

Go supports the use of a switch statement inside another switch statement, commonly referred to as a nested switch.

```go
outer := "go"
    inner := 2

    switch outer {
    case "go":
        fmt.Println("In the outer switch - go")
        switch inner {
        case 1:
            fmt.Println("In the inner switch - case 1")
        case 2:
            fmt.Println("In the inner switch - case 2")
        default:
            fmt.Println("In the inner switch - no case matched")
        }
    case "python":
        fmt.Println("In the outer switch - python")
    default:
        fmt.Println("In the outer switch - no case matched")
    }
```

In this example, there's an outer switch checking the value of the variable `outer`, and an inner switch checking the value of `inner`. Depending on the values of these variables, different cases will be executed. For this specific example, it will print:

```unix
    In the outer switch - go
    In the inner switch - case 2
```

Nested switches can be as deep as you need, though deep nesting can sometimes make the code harder to read and maintain. It's often a good idea to consider if there's a clearer way to write the logic, possibly by breaking it into separate functions or using other control structures.

# Exercice Instructions

In this exercise we will simulate the first turn of a [Blackjack](https://en.wikipedia.org/wiki/Blackjack) game.

You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string such as "ace", "king", "three", "two", etc. The values of each card are:

| card  | value | card    | value |
| :---: | :---: | :-----: | :---: |
|  ace  |  11   | eight   |   8   |
|  two  |   2   | nine    |   9   |
| three |   3   |  ten    |  10   |
| four  |   4   | jack    |  10   |
| five  |   5   | queen   |  10   |
|  six  |   6   | king    |  10   |
| seven |   7   | *other* |   0   |

**Note**: Commonly, aces can take the value of 1 or 11 but for simplicity we will assume that they can only take the value of 11.

Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

- Stand (S)
- Hit (H)
- Split (P)
- Automatically win (W)

Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

- If you have a pair of aces you must always split them.
- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
- If your cards sum up to a value within the range [17, 20] you should always stand.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
- If your cards sum up to 11 or lower you should always hit.

## 1. Calculate the value of any given card.

Implement a function to calculate the numerical value of a card:

```go
value := ParseCard("ace")
fmt.Println(value)
// Output: 11
```

## 2. Implement the decision logic for the first turn.

Write a function that implements the decision logic as described above:

```go
func FirstTurn(card1, card2, dealerCard string) string
```

Here are some examples for the expected outcomes:

```go
FirstTurn("ace", "ace", "jack") == "P"
FirstTurn("ace", "king", "ace") == "S"
FirstTurn("five", "queen", "ace") == "H"
```
