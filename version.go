package goversion

import (
	"fmt"
)

// Compare 2 version formats for which is greater.
// Valid format is as follows: [v]Major.Minor.Build.Version[-{alpha,beta}]
// For example: v0.2.6, 2.7.9-alpha.3, 9.2-beta, 0.2-rc.1
// Returns: 0 if equals, 1 if the first version arg is greater, 2 if the second.
func compareVersions(v1, v2 string) int {

}
