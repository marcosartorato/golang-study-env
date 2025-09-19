package main

import (
	"fmt"
	"reflect"
	"time"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom returns a string representation of the reflected value.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%d", v.Uint())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	case reflect.String:
		return fmt.Sprintf("%q", v.String())
	case reflect.Chan, reflect.Func, reflect.Slice, reflect.Map:
		return fmt.Sprintf("%s 0x%x", v.Type(), v.Pointer())
	case reflect.Struct, reflect.Array:
		return fmt.Sprintf("%s %v", v.Type(), v.Interface())
	case reflect.Interface:
		// The empty interface value is intercepted by reflect.Invalid.
		return formatAtom(v.Elem())
	case reflect.Ptr:
		if v.IsNil() {
			return fmt.Sprintf("%s nil", v.Type())
		}
		return fmt.Sprintf("%s 0x%x", v.Type(), v.Pointer())
	default:
		return v.Type().String() + " value"
	}
}

func main() {
	vIn64, vUint64 := int64(42), uint64(42)
	fmt.Printf("Print int64 %d using reflect functions: %s\n", vIn64, Any(vIn64))
	fmt.Printf("Print uint64 %d using reflect functions: %s\n", vUint64, Any(vUint64))

	vTime := 1 * time.Microsecond
	fmt.Printf("Print time.Duration %s using reflect functions: %s\n", vTime, Any(vTime))

	fmt.Printf("Print string %q using reflect functions: %s\n", "hello", Any("hello"))

	vBool := true
	fmt.Printf("Print bool %t using reflect functions: %s\n", vBool, Any(vBool))

	vCh := make(chan struct{})
	fmt.Printf("Print chan %v using reflect functions: %s\n", vCh, Any(vCh))

	vStruct := struct {
		Name string
		Age  int
	}{Name: "Alice", Age: 30}
	fmt.Printf("Print struct %v using reflect functions: %s\n", vStruct, Any(vStruct))

	var vNilInterface interface{}
	fmt.Printf("Print nil interface using reflect functions: %s\n", Any(vNilInterface))

	var vInterface interface{} = vStruct
	fmt.Printf("Print interface containing struct %v using reflect functions: %s\n", vInterface, Any(vInterface))

	var vNilPtr *int
	fmt.Printf("Print nil pointer using reflect functions: %s\n", Any(vNilPtr))

	vPtr := &vIn64
	fmt.Printf("Print pointer %p using reflect functions: %s\n", vPtr, Any(vPtr))
}
