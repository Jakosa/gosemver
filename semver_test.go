package gosemver

import (
	"testing"
)

func TestCompare(t *testing.T) {

}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compare("v3.0.4-rc.1", "v3.0.5")
	}
}
