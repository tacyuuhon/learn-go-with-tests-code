package reflection

import (
	"reflect"
)

// Walk func
func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
