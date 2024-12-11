package types

type Either[T any] struct {
	Left  error
	Right T
}
