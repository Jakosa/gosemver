package gosemver

import (
	"strconv"
	"strings"
)

// Compare 2 version formats for which is greater.
// For more information see http://semver.org
// Returns: 0 if equals, 1 if the first version arg is greater, 2 if the second.
func Compare(v1, v2 string) int {
	if strings.Contains(v1, "-") {
		strings.Replace(v1, "-", ".", -1)
	}
	if strings.Contains(v2, "-") {
		strings.Replace(v2, "-", ".", -1)
	}
	ver1 := strings.Split(v1, ".")
	ver2 := strings.Split(v2, ".")

	if ver1[0] == ver2[0] {
		if ver1[1] == ver2[1] && len(ver1) > 1 && len(ver2) > 1 {
			if ver1[2] == ver2[2] && len(ver1) > 2 && len(ver2) > 2 {
				return 0
			} else if ver1[2] > ver2[2] {
				return 1
			} else {
				return 2
			}
		} else if ver1[1] > ver2[1] {
			return 1
		} else {
			return 2
		}
	} else if ver1[0] > ver2[0] {
		return 1
	} else {
		return 2
	}
}
