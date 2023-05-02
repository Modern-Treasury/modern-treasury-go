package field

import (
	"fmt"
)

type FieldLike interface {
	field()
}

type Field[T any] struct {
	Value   T
	Null    bool
	Present bool
	Raw     any
}

func (f Field[T]) field() {}
func (f Field[T]) String() string {
	if s, ok := any(f.Value).(fmt.Stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%#v", f.Value)
}
