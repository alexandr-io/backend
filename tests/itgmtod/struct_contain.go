package itgmtod

import (
	"fmt"
	"reflect"
)

func isZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// StructContain compare 2 struct interfaces. If the non-empty data of the y interface are the same as the one in
// the x interface, true is returned. Otherwise false
func StructContain(x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	//if v1.Type() != v2.Type() { // Check that the 2 interfaces are the same struct
	//	fmt.Printf("Expected type: %s != from response type: %s", v2.Type().Name(), v1.Type().Name())
	//	return false
	//}
	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	for i := 0; i < v1.NumField(); i++ { // Go through all the fields of the struct
		val1 := v1.Field(i).Interface()
		val2 := v2.Field(i).Interface()
		if v1.Field(i).Kind() == reflect.Ptr {
			val1 = v1.Field(i).Elem().Interface()
		}
		if v2.Field(i).Kind() == reflect.Ptr {
			val2 = v2.Field(i).Elem().Interface()
		}
		if !isZeroOfUnderlyingType(val2) { // Check that the second struct field is not empty
			if val1 != val2 { // Check if field is equal in two structs
				fmt.Printf("Field %s with value %v is not eqal to field %s with value %v\n",
					v1.Type().Field(i).Name,
					val1,
					v2.Type().Field(i).Name,
					val2)
				return false
			}
		}
	}
	return true
}
