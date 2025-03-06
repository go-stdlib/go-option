// Package option contains functionality for managing
// functional options for fluent instance creation.
package option

// Option defines a functional option for type `T`
// that mutates the instance.
type Option[T any] func(t T) error

// Options is a slice of `Option` options.
type Options[T any] []Option[T]

// WithDefaults prepends defaults to the options slice so they are applied first.
func WithDefaults[T any](options Options[T], defaults ...Option[T]) Options[T] {
	return append(defaults, options...)
}

// Apply applies all functional options for type `T`
// and returns the value or error if any fail to apply.
func Apply[T any](t *T, options ...Option[*T]) (*T, error) {
	for _, o := range options {
		if err := o(t); err != nil {
			return t, err
		}
	}
	return t, nil
}

// New creates a zero value for type `T`, applies all options,
// and returns the value or error if any fail to apply.
func New[T any](options ...Option[*T]) (*T, error) {
	return Apply(new(T), options...)
}
