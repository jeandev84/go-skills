package main

import (
	"fmt"
)

const ERR_CODE_LOCKED = 1
const ERR_CODE_CLOSED = 2

// Error custom
type AppError struct {
	Code int
	Err  error
}

// Implementation de l'interface error de Go
func (e *AppError) Error() string {
	return fmt.Sprintf("code=%v, error=%v", e.Code, e.Err)
}

func process(input int) error {
	if input == 0 {

		return &AppError{
			Code: ERR_CODE_LOCKED,
			Err:  fmt.Errorf("resource is locked"),
		}

	} else if input == 1 {

		return &AppError{
			Code: ERR_CODE_CLOSED,
			Err:  fmt.Errorf("resource is closed"),
		}
	}

	return nil
}

// Entry point
func main() {

	err := process(1)

	if err != nil {

		fmt.Println(err)

		re, ok := err.(*AppError)
		if ok {
			switch re.Code {
			case ERR_CODE_LOCKED:
				fmt.Println("Try again later")
			case ERR_CODE_CLOSED:
				fmt.Println("Resource has to be opened again")
			}
		}
	}
}

/*
code=2, error=resource is closed
Resource has to be opened again
*/
