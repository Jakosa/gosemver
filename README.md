# gosemver [![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/Jakosa/gosemver/trend.png)](https://bitdeli.com/free "Bitdeli Badge")
## LICENCE: GNU GPL v3

### Version comparing library in go (using semantic versioning standard http://semver.org)

Compare 2 version formats for which is greater.<br>
For valid version examples and more information see http://semver.org<br>
Returns: 0 if equals, 1 if the first version arg is greater, 2 if the second, -1 if problem occured.

### Usage

```go
package main

import (
	"fmt"
	"github.com/Jackneill/gosemver"
)

func main() {
	fmt.Println(gosemver.Compare("v0.3.6", "0.4")) // prints 2
}
```
