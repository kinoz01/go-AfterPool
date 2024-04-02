# About

In Go, we can compare values using the equality and ordering operators.
We have 6 of these operators:

| Comparison        | Operator  |
| ------------------| --------- |
| equal             | `==`      |
| not equal         | `!=`      |
| less              | `<`       |
| less or equal     | `<=`      |
| greater           | `>`       |
| greater or equal  | `>=`      |

The result of a comparison is always a boolean value:

```go
a := 3

a != 4 // true
a > 5  // false
```

## Equality operators

The equality operators `==` and `!=` check whether a value is equal to another or not, respectively.
We can use the equality operators on the majority of types in Go. Here are some common examples:

```go
2 == 3 // false, integer comparison

2.1 != 2.2 // true, float comparison 

"hello" == "hello" // true, string comparison

'a' != 'b' // true, rune/byte comparison
```

## Ordering operators

Ordering operators check if one value is greater than (`>`), greater or equal to (`>=`), less than (`<`) and less or equal to (`<=`) to another value.
This kind of comparison is available for numbers, bytes, runes and strings.
When comparing strings, the dictionary order (also known as lexicographic order) is followed.

```go
2 > 3 // false, integer comparison

1.2 < 1.3 // true, float comparison

'a' > 'b' // false, rune/byte comparison

"Hello" < "World" // true, string comparison
```

For more details visit [this arcticle][Article].

[Article]: https://medium.com/golangspec/comparison-operators-in-go-910d9d788ec0

# If statements

Conditionals in Go are similar to conditionals in other languages.
The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`.
Conditionals are often used as flow control mechanisms to check for various conditions.
For checking a particular case an [`if` statement][if_statement] can be used, which executes its code if the underlying condition is `true` like this:

```go
var value string

if value == "val" {
    return "was val"
}
```

In scenarios involving more than one case many `if` statements can be chained together using the `else if` and `else` statements.

```go
var number int
result := "This number is "

if number > 0 {
    result += "positive"
} else if number < 0 {
    result += "negative"
} else {
    result += "zero"
}
```

However, it is convention to avoid `else` statements as Go promotes early returns:

```go
func getVal(connected bool) int {
    // The exceptional case should be in the `if` statemeent.
    // In this case being `connected` is the default, `readLocalVal` the fallback.
    if !connected {
        // using an `early return` to remove the need for an `else` case
        return readLocalVal()
    }

    return readVal()
}
```

If statements can also include a short initialization statement that can be used to initialize one or more variables for the if statement.
For example:

```go
num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} else {
    fmt.Println(num)
}
// Output: 14
```

> Note: just like a for loop any variables created in the initialization statement go out of scope after the end of the if statement.

Coming from other languages one may be tempted to try to use one-line conditionals.
Go does not support ternary operators or one-line conditionals.
This is a purposeful design decision to limit the amount of possibilities, making code simpler and easier to read.

To learn more about this topic it is recommended to check this source: [Go by Example: If/Else][go_by_example_if]

[if_statement]: https://golang.org/ref/spec#If_statements
[go_by_example_if]: https://gobyexample.com/if-else

# Exercice Instructions

In this exercise, you are going to write some code to help you prepare to buy a vehicle.

You have three tasks, one to determine if you need a license, one to help you choose between two vehicles and one to estimate the acceptable price for a used vehicle.

## 1. Determine if you will need a driver's license

Some vehicle kinds require a driver's license to operate them.
Assume only the kinds `"car"` and `"truck"` require a license, everything else can be operated without a license.

Implement the `NeedsLicense(kind)` function that takes the kind of vehicle and returns a boolean indicating whether you need a license for that kind of vehicle.

```go
needLicense := NeedsLicense("car")
// => true

needLicense = NeedsLicense("bike")
// => false

needLicense = NeedsLicense("truck")
// => true
```

## 2. Choose between two potential vehicles to buy

You evaluate your options of available vehicles.
You manage to narrow it down to two options but you need help making the final decision.
For that, implement the function `ChooseVehicle(option1, option2)` that takes two vehicles as arguments and returns a decision that includes the option that comes first in lexicographical order.

```go
vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
// => "Toyota Corolla is clearly the better choice."

ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
// => "Volkswagen Beetle is clearly the better choice."
```

## 3. Calculate an estimation for the price of a used vehicle

Now that you made a decision, you want to make sure you get a fair price at the dealership.
Since you are interested in buying a used vehicle, the price depends on how old the vehicle is.
For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new.
If it is at least 10 years old, it costs 50%.
If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.

Implement the `CalculateResellPrice(originalPrice, age)` function that applies this logic using `if`, `else if` and `else` (there are other ways if you want to practice).
It takes the original price and the age of the vehicle as arguments and returns the estimated price in the dealership.

```go
CalculateResellPrice(1000, 1)
// => 800

CalculateResellPrice(1000, 5)
// => 700

CalculateResellPrice(1000, 15)
// => 500
```

**Note** the return value is a `float64`.