package hooker

import (
	"testing"
)

func TestHookFunc(t *testing.T) {
	var func1 = Hook1(SomeFunc)
	func1(1, 2)
}

func TestHookFunc2(t *testing.T) {
	hooker := NewHooker(SomeFunc, Hook1, Hook1)
	hooker.Run(1, 2)
}
