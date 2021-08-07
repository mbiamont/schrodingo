# Schrodingo
Golang Result Monad library to easily manage errors

## :warning: Warning :warning:

This project requires Go generics which are still in development.

So this project is for experimental usage only.

## Result Monad

This library aims to ease the management of errors in Go using the `Either Monad` which can be either a `Success` (which contains a value), or a `Failure` (which contains an error).

Is it largely inspired by the `Result` class in `Kotlin`.

## Installation

Don't install it until Go generics are a thing.

## Usage

#### Classic error handling 

```go
func DoOtherThing() (int, error){
	return 32, nil
}

func DoSomethingElse(input int) (float32, error){
    return 32.0, nil
}

func DoOneLastThing(input float32) (string, error){
	return "", fmt.Errorf("boom")
}

func DoSomething() error {
	i, err := DoOtherThing()
	
	if err != nil {
		return err
	}
	
	f, err := DoSomethingElse(i)
	
	if err != nil {
		return err
	}
	
	s, err := DoOneLastThing(f)
	
	if err != nil {
		return err
	}
	
	fmt.Println(s)
}

func main() {
	err := DoSomething()
	
	if err != nil {
		panic(err)
	}
}
```

### Using Schrodingo

```go
func DoOtherThing() Result[int]{
	return Success(32)
}

func DoSomethingElse(input int) Result[float32]{
    return Success(32.0)
}

func DoOneLastThing(input float32) Result[string]{
	return Failure(fmt.Errorf("boom"))
}

func DoSomething() error {
	i := DoOtherThing()
	f := ThenDo(i, func(i int) Result[float32] {
		//This function will be called only if DoOtherThing() didn't returned an error
		return DoSomethingElse(i)
	})
	s := ThenDo(f, func(f float32) Result[string] {
		//This function will be called only if DoSomethingElse() didn't returned an error
		return DoOneLastThing(f)
	})
	
	s.OnSuccess(func(s string) {
		//This function will be called only if DoOneLastThing() didn't returned an error
		fmt.Println(s)
	})
	
	return s.ErrorOrNil() //Will return an error if any
}

func main() {
	err := DoSomething()
	
	if err != nil {
		panic(err)
	}
}
```

More infos here: https://melvinbiamont.medium.com/how-generics-can-change-error-handling-in-go-34f47347925a