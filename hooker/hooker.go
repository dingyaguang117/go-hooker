package hooker

// Hook 原型
type Hook[Func any] func(Func) Func

// Hooker 将 Hook 操作封装起来
type Hooker[Func any] struct {
	origin  Func         // 原始函数
	wrapped Func         // 包装后的函数
	hooks   []Hook[Func] // 钩子函数列表
}

func NewHooker[Func any](origin Func, hooks ...Hook[Func]) *Hooker[Func] {
	h := &Hooker[Func]{
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

//func (h *Hooker[Func]) Run(a, b int) int {
//	return h.wrapped(a, b)
//}

func (h *Hooker[Func]) GetWrappedFunc() Func {
	return h.wrapped
}
