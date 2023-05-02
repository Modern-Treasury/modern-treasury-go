package moderntreasury

import (
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
)

func F[T any](value T) field.Field[T]          { return field.Field[T]{Value: value, Present: true} }
func Null[T any]() field.Field[T]              { return field.Field[T]{Null: true, Present: true} }
func Raw[T any](value any) field.Field[T]      { return field.Field[T]{Raw: value, Present: true} }
func Int(value int64) field.Field[int64]       { return F(value) }
func String(str string) field.Field[string]    { return F(str) }
func Float(value float64) field.Field[float64] { return F(value) }
