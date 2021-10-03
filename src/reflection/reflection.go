package reflection

import "reflect"

// interface{} => can pass any type
func walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	// numOfValues := 0                     // how many fields
	// var getField func(int) reflect.Value // how to extract the value (via Field or Index)

	switch val.Kind() {
	case reflect.String:
		// check if the field type is string
		fn(val.String())
	case reflect.Struct:
		// for i := 0; i < val.NumField(); i++ {
		// 	field := val.Field(i)
		// 	walk(field.Interface(), fn)
		// }

		// NumField: return number of fields in the value
		// numOfValues = val.NumField()
		// getField = val.Field

		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		// for i := 0; i < val.Len(); i++ {
		// 	walk(val.Index(i).Interface(), fn)
		// }

		// numOfValues = val.Len()
		// getField = val.Index
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		// calls the function with input nil
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

	// for i := 0; i < numOfValues; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }
}

func getValue(x interface{}) reflect.Value {
	// inspect x
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		// extract the underlying value from the pointer
		val = val.Elem()
	}

	return val
}
