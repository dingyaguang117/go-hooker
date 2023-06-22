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

func Hook1(next Func) Func {
	return func(a, b int) int {
		fmt.Println("----- Hook1 before ----")
		c := next(a, b)
		fmt.Println("----- Hook1 after ----")
		return c
	}
}

func Hook2(next Func) Func {
	return func(a, b int) int {
		fmt.Println("----- Hook2 before ----")
		c := next(a, b)
		fmt.Println("----- Hook2 after ----")
		return c
	}
}

func TestHookFunc(t *testing.T) {
	hooker := NewHooker[Func](SomeFunc, Hook1, Hook2).AddHook(Hook2)
	hooker.GetWrappedFunc()(1, 2)
}
