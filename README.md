# random-string
Generate random string in Golang

## Install
```bash
go get TheYahya/random-string
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/theyahya/random-string"
)

func main() {
    randomString := randomstring.Generate(10)
	fmt.Println(randomString)
}
```

## License
Licensed under the [MIT License](https://github.com/TheYahya/random-string/blob/main/README.md).