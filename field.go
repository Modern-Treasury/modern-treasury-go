package moderntreasury

import (
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
)

func F[T any](value T) param.Field[T]          { return param.Field[T]{Value: value, Present: true} }
func Null[T any]() param.Field[T]              { return param.Field[T]{Null: true, Present: true} }
func Raw[T any](value any) param.Field[T]      { return param.Field[T]{Raw: value, Present: true} }
func Int(value int64) param.Field[int64]       { return F(value) }
func String(str string) param.Field[string]    { return F(str) }
func Float(value float64) param.Field[float64] { return F(value) }
