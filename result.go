package schrodingo

import "fmt"

type any struct {
}

type Void interface {
}

/* --------------- */


func DoSomething() Result[int] {
	return Success(3)
}

func DoSomethingElse(input int) Result[float64] {
	return Success(3.141592)
}

func DoOneLastThing(input float64) Result[string] {
	return Success("3.141592")
}

func main() {
	i := DoSomething()
	f := Map(i, func(i int) Result[float64] {
		return DoSomethingElse(i)
	})
	s := Map(f, func(f float64) Result[string] {
		return DoOneLastThing(f)
	})

	s.Fold(func(s string) {
		fmt.Println(s)
	}, func(err error) {
		panic(err)
	})
}

/* --------------- */

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

func (r Result[T]) get() *T {
	if v, ok := r.v.(T); ok {
		return &v
	}

	return nil
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