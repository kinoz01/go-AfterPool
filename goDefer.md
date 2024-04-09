In Go (Golang), the `defer` statement is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. `defer` is often used where you would use `ensure` or `finally` in other languages. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

## Basic Syntax

The basic syntax of a `defer` statement is:

```go
defer functionCall()
```

## Characteristics of `defer`

1. **Deferred function calls are pushed onto a stack**. When a function returns, its deferred calls are executed in last-in, first-out order.
2. **Arguments evaluation**: The arguments of the deferred function are evaluated when the `defer` statement is executed, not when the function call is executed.
3. **Used for resource cleanup**: It's often used for resources that need to be released regardless of which path a function takes to return. Common examples include closing files, unlocking mutexes, or closing network connections.

## Examples

**Example 1: Deferring a Simple Function Call**

```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// output
// hello
// world
```

In this example, `fmt.Println("world")` is deferred, so it executes after the surrounding `function (main)` has finished executing, hence "hello" is printed before "world".

**Example 2: Deferring a Function Call to Close a File**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("filename.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Operations on the file
    fmt.Println("File opened successfully")
}
```

Here, `defer f.Close()` ensures that the file is closed, regardless of whether the subsequent operations succeed or fail.

**Example 3: Deferred Calls Order**

```go
package main

import "fmt"

func main() {
    fmt.Println("start")
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("end")
}
```

```unix
start
end
2
1
0
```
This demonstrates how deferred calls are stacked and then executed in last-in, first-out order when the surrounding function returns.

## Usage Cases

1. **Resource Cleanup**: It's common to defer calls to Close methods on resources like files and network connections.
2. **Unlocking Mutexes**: After locking a mutex, you can defer its unlocking to ensure it gets unlocked as the function returns.
3. **Printing Final Messages**: For example, printing a function's execution time by deferring the end time calculation and printing at the start of a function.
4. **More Complex Cleanup Logic**: You can defer anonymous functions that perform complex cleanup tasks, allowing you to use the surrounding function's variables for cleanup logic.