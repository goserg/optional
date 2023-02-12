package optional

import "fmt"

type Optional[T any] struct {
	value T
	ok    bool
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func New[T any](value T) Optional[T] {
	return Optional[T]{
		value: value,
		ok:    true,
	}
}

func (o Optional[T]) Get() (T, bool) {
	return o.value, o.ok
}

func (o Optional[T]) Must() T {
	if !o.ok {
		panic("optional value is empty")
	}
	return o.value
}

func (o Optional[T]) GetOrElse(e T) T {
	if !o.ok {
		return e
	}
	return o.value
}

func (o Optional[T]) String() string {
	if !o.ok {
		return fmt.Sprintf("None(%T)[]", o.value)
	}
	value := o.Must()
	if stringer, ok := interface{}(value).(fmt.Stringer); ok {
		return fmt.Sprintf("Some(%T)[%s]", value, stringer)
	}
	return fmt.Sprintf("Some(%T)[%v]", value, value)
}
