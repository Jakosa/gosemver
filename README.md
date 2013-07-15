# gosemver
### LICENCE: GNU GPL 3

### Version comparing library in go (using semantic versioning standard http://semver.org)

Compare 2 version formats for which is greater.<br>
For valid version examples and more information see http://semver.org<br>
Returns: 2 values. 
* First: 0 if equals, 1 if the first version arg is greater, 2 if the second.
* Second: error if any, otherwise nil.

### Usage

```go
package main

import (
	"fmt"
	"github.com/Jackneill/gosemver"
)

func main() {
	result, nil := gosemver.Compare("v0.3.6", "0.4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}```