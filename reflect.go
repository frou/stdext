package stdext

import "reflect"

// CollapseSingleElementSlice returns x as-is, unless x is a slice containing a
// single element, in which case that element itself is returned.
func CollapseSingleElementSlice(x interface{}) interface{} {
	reflectedVal := reflect.ValueOf(x)
	if reflectedVal.Kind() == reflect.Slice && reflectedVal.Len() == 1 {
		return reflectedVal.Index(0).Interface()
	}
	return x
}
