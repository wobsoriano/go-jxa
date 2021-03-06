# go-jxa

Go module for running JavaScript for Automation (JXA). Requires macOS 10.10 or later.

## Install

```bash
$ go get github.com/wobsoriano/go-jxa
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	jxa "github.com/wobsoriano/go-jxa"
)

func main() {
	code := "Application('System Events').appearancePreferences.darkMode()"
	v, err := jxa.RunJXA(code)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Is dark mode: %s", v)
}
```

## License

MIT
