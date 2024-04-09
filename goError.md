# Errors Handling

In Go (Golang), handling and reporting errors effectively is crucial for creating robust applications. Go takes a somewhat unique approach to error handling compared to many other programming languages. It does not use exceptions; instead, errors are values that can be returned and handled explicitly. Here are some of the most common ways to return and report errors in Go, including examples using the `fmt`, `os`, and `errors` packages, as well as other methods:

## 1. Returning an Error

One of the most basic ways to handle errors in Go is by returning an error from a function. If a function can fail, it should return an error as its last return value.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## 2. Printing an Error to `Stdout` or `Stderr` Using `fmt`

The `fmt` package provides functions to format and print text, including errors.

**Printing to stdout:**

```go
if result, err := divide(10, 0); err != nil {
    fmt.Println("Error:", err)
}
```

**Printing to stdout:**

```go
if result, err := divide(10, 0); err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
}
```

## 3. Creating Custom Errors with `errors.New`

You can create custom error messages using the `errors.New` function. It's a straightforward approach when you need a simple error without formatting.

```go
import "errors"

func myFunction() error {
    return errors.New("a simple error")
}
```

## 4. Formatting Errors with `fmt.Errorf`

When you need to include dynamic data in your error messages, `fmt.Errorf` is very handy. It allows you to format the error message similarly to `fmt.Sprintf`.

```go
func anotherFunction() error {
    return fmt.Errorf("error occurred at %v", time.Now())
}
```

## 5. Using `os.Exit` to Exit the Program

If an error is critical enough that the program should not continue executing, you can use `os.Exit` to exit the program immediately. Note that `os.Exit` does not call defer statements.

```go
if err := criticalFunction(); err != nil {
    fmt.Fprintln(os.Stderr, "Critical error occurred:", err)
    os.Exit(1)
}
```

## 6. Panicking (Not Recommended for Common Error Handling)

Panicking is a way to handle errors that should never occur under normal conditions. It's similar to throwing an exception in other languages. However, using panic for regular error handling is discouraged; it's primarily meant for situations where the program cannot recover.

```go
func riskyFunction() {
    if badThingHappened {
        panic("something really bad happened")
    }
}
```

## 7. Returning Custom Error Types

You can define custom error types by implementing the error interface, which requires a method `Error() string`.

```go
type MyError struct {
    Message string
    Code    int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func triggerError() error {
    return &MyError{"Something happened", 400}
}
```

# Other Examples

## 1. `log` Package

The `log` package can be used for logging errors, with the capability to prepend logs with a timestamp or additional information automatically. It's more sophisticated than simple `fmt.Print` statements and can be directed to various outputs, including files or `stderr`.

```go
import "log"

func someFunction() {
    err := anotherFunction()
    if err != nil {
        log.Println("Error:", err)
    }
}

// Setting log output to stderr (default is stderr, but this is for demonstration)
log.SetOutput(os.Stderr)
```

## 2. `io/ioutil` and `io` Packages

While not directly related to error handling, the `io` and `io/ioutil` packages (with `ioutil` now being recommended to use through `os` and `io` as of Go 1.16) often work with errors returned from IO operations, enabling handling of IO-related errors.

```go
import (
    "io"
    "os"
)

func copyFileContents(src, dst string) error {
    inFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer inFile.Close()

    outFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, inFile)
    if err != nil {
        return err
    }

    return nil
}
```

## 3. `net` Package

Errors from the `net` package are common when dealing with network operations, such as opening connections, reading from or writing to connections, etc. Handling these errors allows for robust networking applications.

```go
import (
    "net"
    "fmt"
)

func connectToServer(addr string) {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()
    // Use the connection
}
```

## 4. `syscall` Package

The `syscall` package provides an interface to the low-level operating system primitives. Errors from system calls are common when dealing with operations such as file manipulation, process management, etc., that require direct system calls.

```go
import (
    "fmt"
    "syscall"
)

func setFileReadOnly(path string) error {
    err := syscall.Chmod(path, 0400) // Set file as read-only
    if err != nil {
        return fmt.Errorf("failed to set file as read-only: %w", err)
    }
    return nil
}
```

## 5. `http` Package

When building web applications, handling errors from the http package is crucial for responding to HTTP requests properly.

```go
import (
    "net/http"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    _, err := w.Write([]byte("Hello, world!"))
    if err != nil {
        log.Printf("Error writing response: %v", err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Each of these packages provides its own set of functions and methods that can return errors. Properly handling these errors is key to writing reliable and maintainable Go programs.