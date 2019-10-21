package main

import (
	"errors"
	"fmt"

	"github.com/henzai/exc/v0"
)

const (
	// NotFound is represented
	NotFound exc.Kind = "not found"
	// Internal b
	Internal exc.Kind = "internal error"
)

func main() {
	got, err := controller(1)
	if err != nil {
		var baseErr *exc.Error
		if ok := errors.As(err, &baseErr); ok {
			switch k := baseErr.LastKind(); k {
			case NotFound:
				fmt.Println("gggggggggggggg")
			default:
				fmt.Println(k)
			}
		}
		fmt.Printf("%+v", err)
		return
	}
	fmt.Println(got)
}

func controller(id int) (string, error) {
	got, err := usecase(id)
	if err != nil {
		return "", exc.Wrap(err, exc.WithMessage("error controller"))
	}
	return got, nil
}

func usecase(id int) (string, error) {
	got, err := repository(id)
	if err != nil {
		return "", exc.Wrap(err, exc.WithMessagef("usecase %v", 1), exc.WithKind(Internal))
	}
	return got, nil
}

func repository(id int) (string, error) {
	return "", exc.Wrap(errors.New("origin"), exc.WithMessage("repository"), exc.WithKind(Internal))
}
