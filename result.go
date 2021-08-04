package schrodingo

type Result[T any] struct {
	v   interface{}
	err error
}

func Success[T any](value T) Result[T] {
	output := new(Result[T])
	output.v = value

	return *output
}

func Failure[T any](e error) Result[T] {
	output := new(Result[T])
	output.err = e

	return *output
}

func (r Result[T]) IsFailure() bool {
	return r.v == nil
}

func (r Result[T]) IsSuccess() bool {
	return r.v != nil
}

func (r Result[T]) GetOrElse(defaultValue T) T {
	if r.IsSuccess() {
		if v, ok := r.v.(T); ok {
			return v
		}
	}
	return defaultValue
}

func (r Result[T]) GetOrNil() *T {
	if r.IsSuccess() {
		if v, ok := r.v.(T); ok {
			return &v
		}
	}
	return nil
}

func Map[T any, R any](result Result[T], mapper func(T) Result[R]) Result[R] {
	if result.IsSuccess() {
		if v, ok := result.v.(T); ok {
			return mapper(v)
		}
	}
	return Failure[R](result.err)
}

func (r Result[T]) OnSuccess(onSuccess func(T)) Result[T] {
	if r.IsSuccess() {
		if v, ok := r.v.(T); ok {
			onSuccess(v)
		}
	}

	return r
}

func (r Result[T]) OnFailure(onFailure func(error)) Result[T] {
	if r.IsFailure() {
		onFailure(r.err)
	}

	return r
}

func (r Result[T]) Fold(onSuccess func(T), onFailure func(error)) Result[T] {
	if r.IsSuccess() {
		if v, ok := r.v.(T); ok {
			onSuccess(v)
		}
	} else {
		onFailure(r.err)
	}

	return r
}