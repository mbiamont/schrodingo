package schrodingo

type Result[T any] interface {
	IsSuccess() bool
	IsFailure() bool

	GetOrElse(defaultValue T) T
	GetOrNil() *T
	ErrorOrNil() error

	OnSuccess(onSuccess func(T)) Result[T]
	OnFailure(onFailure func(error)) Result[T]
	Fold(onSuccess func(T), onFailure func(error)) Result[T]
}

func Success[T any](value T) Result[T] {
	output := new(success[T])
	output.v = value

	return *output
}

func Failure[T any](e error) Result[T] {
	output := new(failure[T])
	output.err = e

	return *output
}

func ThenDo[T any, R any](result Result[T], nextFunction func(T) Result[R]) Result[R] {
	if result.IsSuccess() {
		return nextFunction(result.(success[T]).v)
	}
	return Failure[R](result.ErrorOrNil())
}