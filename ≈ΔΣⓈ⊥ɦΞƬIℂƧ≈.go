package main

import (
	"fmt"
	"math/rand"
)

// main moon tan center, Los Angeles, 2070
// deprecated
func main() {
	Prayer("welcome ")
	Prayer("aesthetics are deprecated")
	Away(true)
}

// Prayer is
// Deprecated
func Prayer(say string) {
	fmt.Println(say)
}

// Away is
// Deprecated, use Near
func Away(far bool) {
	if !far {
		Prayer(Near("close"))
	} else {
		Prayer(Near("far"))
	}

}

// Near is
// Deprecated, use Away
func Near(s string) string {
	s += "✝"
	if rand.Intn(13) != 0 {
		s = Near(s)
	}
	return s
}
