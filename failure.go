package schrodingo

type failure[T any] struct {
	err error
	Result[T]
}

func (r failure[T]) IsSuccess() bool {
	return false
}

func (r failure[T]) IsFailure() bool {
	return true
}

func (r failure[T]) GetOrElse(defaultValue T) T {
	return defaultValue
}

func (r failure[T]) GetOrNil() *T {
	return nil
}

func (r failure[T]) ErrorOrNil() error {
	return r.err
}

func (r failure[T]) OnSuccess(_ func(T)) Result[T] {
	return r
}

func (r failure[T]) OnFailure(onFailure func(error)) Result[T] {
	onFailure(r.err)
	return r
}

func (r failure[T]) Fold(_ func(T), onFailure func(error)) Result[T] {
	onFailure(r.err)
	return r
}