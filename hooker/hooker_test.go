package hooker

import (
	"fmt"
	"testing"
)

// Func 原始函数原型
type Func func(a, b int) int

// SomeFunc 原始函数定义
func SomeFunc(a, b int) int {
	fmt.Printf("%d + %d\n", a, b)
	return a + b
}

// Hook1 Hook
var Hook1 Hook[Func] = func(f Func) Func {
	return func(a, b int) int {
		fmt.Println("----- before ----")
		c := f(a, b)
		fmt.Println("----- after ----")
		return c
	}
}

func TestHookFunc(t *testing.T) {
	var func1 = Hook1(SomeFunc)
	func1(1, 2)
}

func TestHookFunc2(t *testing.T) {
	hooker := NewHooker[Func](SomeFunc, Hook1, Hook1)
	hooker.GetWrappedFunc()(1, 2)
}
