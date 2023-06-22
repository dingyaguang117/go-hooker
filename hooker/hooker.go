package hooker

type Hook[Func any] func(Func) Func

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

func (h *Hooker[Func]) AddHook(hook Hook[Func]) *Hooker[Func] {
	h.hooks = append(h.hooks, hook)
	h.wrapped = hook(h.wrapped)
	return h
}

func (h *Hooker[Func]) GetOrigin() Func {
	return h.origin
}

func (h *Hooker[Func]) GetWrapped() Func {
	return h.wrapped
}

func (h *Hooker[Func]) GetHooks() []Hook[Func] {
	return h.hooks
}
