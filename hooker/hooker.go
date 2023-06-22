package hooker

import (
	"fmt"
)

// Func 原始函数原型
type Func func(a, b int) int

// Hook 原型
type Hook func(Func) Func

// SomeFunc 原始函数定义
func SomeFunc(a, b int) int {
	fmt.Printf("%d + %d\n", a, b)
	return a + b
}

// Hook1 Hook
var Hook1 Hook = func(f Func) Func {
	return func(a, b int) int {
		fmt.Println("----- before ----")
		c := f(a, b)
		fmt.Println("----- after ----")
		return c
	}
}

// Hooker 将 Hook 操作封装起来
type Hooker struct {
	origin  Func   // 原始函数
	wrapped Func   // 包装后的函数
	hooks   []Hook // 钩子函数列表
}

func NewHooker(origin Func, hooks ...Hook) *Hooker {
	h := &Hooker{
		origin: origin,
		hooks:  hooks,
	}
	wrapped := origin
	for _, hook := range h.hooks {
		wrapped = hook(wrapped)
	}
	h.wrapped = wrapped
	return h
}

func (h *Hooker) Run(a, b int) int {
	return h.wrapped(a, b)
}
