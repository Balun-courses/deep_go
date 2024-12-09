package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Data struct {
	age     int8
	address string
}

func getFieldOffset(instance interface{}, fieldName string) (uintptr, error) {
	instanceValue := reflect.ValueOf(instance)
	if instanceValue.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("the first parameter must be a pointer to a struct")
	}

	instanceType := instanceValue.Type().Elem()
	field, found := instanceType.FieldByName(fieldName)
	if !found {
		return 0, fmt.Errorf("field '%s' not found in the struct", fieldName)
	}

	return field.Offset, nil
}

func AssignPrivateField[T any, V any](src *T, field string, value V) error {
	offset, err := getFieldOffset(src, field)
	if err != nil {
		return fmt.Errorf("get offset: %w", err)
	}

	fieldPtr := (*V)(unsafe.Add(unsafe.Pointer(src), offset))
	*fieldPtr = value
	return nil
}

func ReadPrivateField[T any, V any](src *T, field string) (V, error) {
	var output V

	offset, err := getFieldOffset(src, field)
	if err != nil {
		return output, fmt.Errorf("get offset: %w", err)
	}

	output = *(*V)(unsafe.Add(unsafe.Pointer(src), offset))
	return output, nil
}

func main() {
	data := &Data{}
	_ = AssignPrivateField(data, "address", "Moscow")
	fmt.Println(data)
}
