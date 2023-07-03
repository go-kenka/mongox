package utils

import "reflect"

func IsZero[T any](v T) bool {
	return reflect.ValueOf(&v).Elem().IsZero()
}

func Nil[T any]() T {
	var zero T
	return zero
}
