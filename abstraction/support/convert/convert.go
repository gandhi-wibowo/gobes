package convert

// Default returns the first non-zero value.
// If all values are zero, return the zero value.
//
//	Default("", "foo") // "foo"
//	Default("bar", "foo") // "bar"
//	Default("", "", "foo") // "foo"
func Default[T comparable](values ...T) T {
	var zero T
	for _, value := range values {
		if value != zero {
			return value
		}
	}
	return zero
}

// Pointer returns a pointer to the value.
//
//	Pointer("foo") // *string("foo")
//	Pointer(1) // *int(1)
func Pointer[T any](value T) *T {
	return &value
}
