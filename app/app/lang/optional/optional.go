package optional

type Optional[Value any] struct {
	value  Value
	exists bool
}

func (o Optional[Value]) Get() (Value, bool) {
	return o.value, o.exists
}

func (o Optional[Value]) Or(deflt Value) Value {
	if o.exists {
		return o.value
	}
	return deflt
}

func (o Optional[Value]) AsPointer() *Value {
	if o.exists {
		return &o.value
	}
	return nil
}

func Map[V any, R any](opt Optional[V], f func(V) R) Optional[R] {
	if opt.exists {
		return Of(f(opt.value))
	}
	return Empty[R]()
}

func Empty[Value any]() Optional[Value] {
	var empty Value
	return Optional[Value]{empty, false}
}

func Of[Value any](value Value) Optional[Value] {
	return Optional[Value]{value, true}
}

func OfPointer[Value any](p *Value) Optional[Value] {
	if p == nil {
		return Empty[Value]()
	}
	return Of(*p)
}
