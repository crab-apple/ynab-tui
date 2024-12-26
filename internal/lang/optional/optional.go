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
