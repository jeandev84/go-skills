package main

import (
	"fmt"
)

// entry point program
func main() {

	// day of month [1, 31]
	// exmaple days: 8, 18, 20
	day := 20

	if day < 15 {
		fmt.Printf("We're in in the first half of the month (day=%d)\n", day)
	} else if day == 18 {
		/* fmt.Printf("We're in the special day (day=%d - %d)\n", day, 10) */
		fmt.Printf("We're in the special day (day=%d)\n", day)
	} else {
		fmt.Printf("We're in the second half of the month (day=%d)\n", day)
	}

	// test boolean
	year, month, day := 2009, 11, 10
	fmt.Printf("Date=%d/%d/%d\n", day, month, year)

	if year == 2009 && month == 11 && day == 10 {
		fmt.Println("This is the first realease of Go!")
	} else if year == 2009 || month == 11 || day == 10 {
		fmt.Println("At least one part is right...:/")
	} else {
		fmt.Println("Just another day...")
	}

	// Ecrire des conditions sur une seule ligne
	// on declare une variable count egal 12, ensuite on ecrit la condition
	// ma variable count sera accessaible que dans mon block if/else (scope de visibilite)
	if count := 12; count > 10 {
		fmt.Printf("We have enough count. got=%d\n", count)
	} else {
		fmt.Printf("Not enough. got=%d\n", count)
	}

	/* fmt.Println(count) : on aura une erreur ici car count n'est visible que dans le block */
}

/*
Println - chaine sure une ligne
Printf  - message formater
%s - chaine de caractere (string)
%d - nombre decimal (float)

if (year == 2009 || month == 11) && day == 10 {
		fmt.Println("At least one part is right...:/")
}
*/
