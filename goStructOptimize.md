# Memory Alignment in Go

When we talk about alignment boundaries in the context of computer memory, we're referring to the way data is organized and accessed in memory. Different types of data (like bool, int16, float32, etc.) require different amounts of memory space, and the hardware prefers accessing this data at certain "aligned" addresses to improve access speed and efficiency.

## Example Struct in Go

Let's break down the example of the exampleStruct in Go to understand alignment:

```go
type exampleStruct struct {
    b bool      // 1 byte, 0
                // 1 byte padding for alignment, 1
    i int16     // 2 bytes, 2-3
    f float32   // 4 bytes, 4-7
}
```

Boolean (bool): A bool in Go occupies 1 byte. Let's say it starts at byte 0 in memory. Since the next element (int16) needs to be aligned on a 2-byte boundary, there's a padding of 1 byte after the bool to ensure the int16 starts at the correct boundary.

Integer 16 bits (int16): int16 needs 2 bytes. According to the alignment rules, it should start at a 2-byte boundary. Thanks to the padding added after the bool, the int16 starts at byte 2, satisfying the alignment requirement.

Floating point 32 bits (float32): float32 needs 4 bytes. The alignment rule for 4-byte values is that they should start at a 4-byte boundary. Since the int16 ends at byte 3, the float32 can start at byte 4, perfectly aligning with the rule.

## Memory Layout Visualization

```CSS
  Memory Address: 0   1   2   3   4   5   6   7
                | b | p | i | i | f | f | f | f |
                  ^   ^   ^       ^
                  |   |   |       |
                bool  |  int16   float32
                   padding (for alignment)
``` 

b: Represents the bool occupying 1 byte.  
p: Represents the padding of 1 byte to align the next data on its boundary.  
i: Represents the int16, which occupies 2 bytes and starts at a 2-byte boundary.  
f: Represents the float32, which occupies 4 bytes and starts at a 4-byte boundary.  

For minimizing internal padding within a struct, leading to more efficient memory usage it is advised to always always lay the fields out from highest to lowest.

If we rearrange the fields in the previous struct from largest to smallest, the struct would look like this:

```go
type exampleStructOptimized struct {
    f float32   // 4 bytes, starts at 0, no padding required for alignment
    i int16     // 2 bytes, starts at 4, aligned on a 2-byte boundary, no padding required
    b bool      // 1 byte, starts at 6, 1 byte of padding might be added at the end for alignment in arrays or larger structs 
}
```

By reordering the fields from largest to smallest, we've eliminated the padding between the fields themselves. This can lead to a more compact structure, saving memory space, especially when many instances of the struct are used.