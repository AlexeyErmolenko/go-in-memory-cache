Go In Memory Cache
================================

Go In Memory Cache helps you to cache different data in memory and manage it.

## Example

```go
package main

import (
  "fmt"
  
  cache "github.com/AlexeyErmolenko/go-in-memory-cache"
)

func main() {
    c := cache.New()
    c.Set("userId", 10)
    fmt.Println(c.Get("userId"))
    c.Delete("userId")
    fmt.Println(c.Get("userId"))
}

```