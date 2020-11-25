package main 

import fmt 


func main()  {
	fmt.Printf("Hello")
}

// Compiler ce programme comme suit
// pour adapter a toute sorte de machine:

env GOOS=linux GOARCH=amd64 go build