package main

import "fmt"

type ErrorCode int

const (
	ERR_CODE_OK ErrorCode = iota
	ERR_CODE_NOT_FOUND
	ERR_CODE_LOCKED
	ERR_CODE_GENERIC
)

func (ec ErrorCode) String() string {
	// ... indique la taille de l'array
	// qui sera égale au nombre d'éléments déclarés
	return [...]string{
		"ok",
		"not found",
		"locked",
		"generic"}[ec]
}

func (ec ErrorCode) IsCritical() bool {
	return ec == ERR_CODE_LOCKED || ec == ERR_CODE_NOT_FOUND
}

func printErrCode(c ErrorCode) {
	fmt.Printf("code=%v, critical=%v, detail=%v\n", c, c.IsCritical(), c.String())
}

func main() {
	code := ERR_CODE_OK
	printErrCode(code)

	printErrCode(ERR_CODE_LOCKED)

	// Attention! ErrorCode reste un int.
	// Cette instruction va compiler mais provoquera un crash
	// lors de l'exécution
	//printErrCode(10)
}
