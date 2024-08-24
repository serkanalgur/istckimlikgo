[![PkgGoDev](https://pkg.go.dev/badge/serkanalgur/istckimlikgo)](https://pkg.go.dev/github.com/serkanalgur/istckimlikgo) [![Go Report Card](https://goreportcard.com/badge/github.com/serkanalgur/istckimlikgo)](https://goreportcard.com/report/github.com/serkanalgur/istckimlikgo) 

# Turkish ID Number Validator
Turkish ID Number Validator for Go

## Installation

```bash
go get -u github.com/serkanalgur/istckimlikgo
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/serkanalgur/istckimlikgo"
    "log"
)

func main(){

    valid, err := istckimlikgo.validate("11111111111")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(valid)
}
```
