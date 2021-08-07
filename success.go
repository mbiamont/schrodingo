package schrodingo

type success[T any] struct {
	v T
	Result[T]
}

func (r success[T]) IsSuccess() bool {
	return true
}

func (r success[T]) IsFailure() bool {
	return false
}

func (r success[T]) GetOrElse(_ T) T {
	return r.v
}

func (r success[T]) GetOrNil() *T {
	return &r.v
}

func (r success[T]) ErrorOrNil() error {
	return nil
}

func (r success[T]) OnSuccess(onSuccess func(T)) Result[T] {
	onSuccess(r.v)
	return r
}

func (r success[T]) OnFailure(_ func(error)) Result[T] {
	return r
}

func (r success[T]) Fold(onSuccess func(T), _ func(error)) Result[T] {
	onSuccess(r.v)
	return r
}