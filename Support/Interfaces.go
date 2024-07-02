package Support

// Implements
// Checks if value implements a specific interface
func Implements[T any](value any) bool {
	_, ok := value.(T)

	return ok
}

// Cast
// Casts a value to a specific interface
func Cast[T any](value any) T {
	cast, ok := value.(T)

	if !ok {
		panic("value can not be cast to T")
	}

	return cast
}
