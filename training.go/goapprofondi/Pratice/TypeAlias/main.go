package main

import (
	"fmt"
)

// type alias
type ErrorCode int

// Groupement de constante
// iota c'est un mot cle de go qui va incrementer la valeur suvante de 1
// en commencant par 0 par default
// ERR_CODE_OK = iota + 1 // ici on le fait commencer par 1
const (
	ERR_CODE_OK ErrorCode = iota
	ERR_CODE_NOT_FOUND
	ERR_CODE_LOCKED
	ERR_CODE_GENERIC
)

// Determine if error code is critical
func (ec ErrorCode) IsCritical() bool {

	return ec == ERR_CODE_LOCKED || ec == ERR_CODE_NOT_FOUND
}

func isValid(ec ErrorCode) bool {
	// voir ci on est dans la table d'enumeration de valeur
}

func (ec ErrorCode) String() string {

	/*
		switch ec {
		case ERR_CODE_NOT_FOUND:
			return "not found"
		}

		un struct anonyme
		[...] la taille de mon tableau correspond
		au 3 petits points
	*/

	return [...]string{
		"ok",
		"not found",
		"locked",
		"generic",
	}[ec]

}

// Print error
func printErrCode(c ErrorCode) {
	fmt.Printf("code=%v, critical=%v, detail=%v\n",
		c,
		c.IsCritical(),
		c)
}

// Entry point
func main() {

	// code := ERR_CODE_OK
	// code := ERR_CODE_LOCKED
	code := ERR_CODE_NOT_FOUND
	printErrCode(code)
}
