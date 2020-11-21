package main

import "fmt"

/*
Syntaxe "For"
for <initialisation>; <condition>; <action de boucle> {
	//code
}

for i := 0; i < 5; i++ {
   fmt.Println(i) // output, 0, 1, 2, 3, 4
}

Syntaxe "While"
for <condition> {
	//code
}

i := 1
for i <= 3 {
	fmt.Println(i) // output: 1, 2, 3
	i = i + 1
}


Syntaxe "Forever"
i := 0
for {
	fmt.Println("Loooop") // output: Loooop, Loooop, Loooop, ...
}

*/

func main() {

	// BOUCLE FOR
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Println("Loop:", i)
		sum += i // sum = sum + i
	}

	fmt.Printf("Sum result=%v\n", sum)

	// BOUCLE WHILE ( Compter d'evenements )
	eventCnt := 0

	for eventCnt < 3 {

		fmt.Println("Retrieving events...")
		eventCnt++

		if eventCnt == 3 {
			fmt.Printf("Got %d notifications, updating Phone notifications\n", eventCnt)
		}
	}

	// BOUCLE INFINIE (Comment sortir d'une boucle)
	i := 0
	for {
		i++

		/*
		 Si le reste de la division entiere par 2 est different de 0, si nombre impair
		 On continue la boucle
		*/
		if i%2 != 0 { // nombre impair
			fmt.Println("Odd looping...")
			continue
		}

		fmt.Println("Looping...")

		/*
		 Sortie de boucle
		*/
		if i >= 10 {
			fmt.Println("Loop end")
			break
		}
	}
}
