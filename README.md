# gocmdln - Command line parsing tools for Go

[![Build Status](https://travis-ci.org/mauricelam/gocmdln.svg?branch=master)](https://travis-ci.org/mauricelam/gocmdln) [![GoDoc](https://godoc.org/github.com/mauricelam/gocmdln/params?status.png)](http://godoc.org/github.com/mauricelam/gocmdln/params)

Get:

```
go get github.com/mauricelam/gocmdln
```

Develop:

```
go get -t github.com/mauricelam/gocmdln
```

-----

## Usage

```go
package main

import (
  "flag"

  "github.com/mauricelam/gocmdln/params"
)

// Usage: sed [OPTION]... <command> [input-files]...
func main() {
  command := params.String("command", /* optional */ true, /* metadata */ nil);
  inputFiles := params.StringList("inputFiles", /* optional */ true, /* metadata */ nil);

  copy := flag.Bool("copy", false, "use copy instead of rename when shuffling files in -i mode")
  
  flag.Parse()
  err := params.Parse(flag.Args())
  if err != nil {
    panic(err)
  }
  
  // Do something
}
```

## Examples

See `examples` directory for more examples

## Contributions

  * Please do TDD
  * All inputs welcome
