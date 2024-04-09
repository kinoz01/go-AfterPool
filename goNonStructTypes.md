# Go Non-Struct Types

In Go, you can define custom types based on existing ones, including both basic and composite types. This feature allows you to create semantically meaningful types, even when they're based on simple underlying data structures. When you define a custom type, you're not just creating an alias; you're creating a new type with the same underlying data structure as the original. This new type can have its own methods, thereby extending its functionality beyond that of the original type.

Let's explore some additional examples where custom types can be particularly useful:

## 1. Enhancing Readability and Type Safety

Custom types can make your code more readable by adding semantic meaning to otherwise generic types. They also enhance type safety by preventing mix-ups between different kinds of data that happen to use the same underlying type.

**Example: Distinct ID Types**

```go
type UserID int
type ProductID int

func NewUser(id UserID) {
    // Create a new user with UserID
}

func NewProduct(id ProductID) {
    // Create a new product with ProductID
}

var uID UserID = 1
var pID ProductID = 1

// NewUser(uID) // Correct
// NewUser(pID) // Compile-time error: cannot use pID (type ProductID) as type UserID
```

Here, even though both `UserID` and `ProductID` are of type `int`, they are considered different types. This prevents you from accidentally using a `ProductID` where a `UserID` is expected, and vice versa.

## 2. Encapsulating Complex Behavior

You can define methods on custom types to encapsulate complex behavior related to the data the type represents.

**Example: Custom Slice Type with Convenience Methods**

```go
package main

type StringSet []string

// Add adds a string to the set, ensuring no duplicates.
func (s *StringSet) Add(str string) {
    for _, existing := range *s {
        if existing == str {
            return
        }
    }
    *s = append(*s, str)
}

// Contains checks if a string is in the set.
func (s *StringSet) Contains(str string) bool {
    for _, existing := range *s {
        if existing == str {
            return true
        }
    }
    return false
}

func main() {
    var mySet StringSet
    mySet.Add("apple")
    mySet.Add("banana")
    fmt.Println(mySet.Contains("apple")) // Output: true
    fmt.Println(mySet.Contains("cherry")) // Output: false
}
```

This example creates a StringSet, a custom type for handling a collection of strings without duplicates. The Add method adds a string to the set only if it's not already present, and the Contains method checks for the presence of a string.
