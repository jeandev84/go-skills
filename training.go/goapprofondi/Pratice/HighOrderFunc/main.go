]package main

import "fmt"

// student struct
type student struct {
	name       string
	grade      int
	attendance int
}

// Filter students
func filter(students []student, f func(student) bool) []student {
	var arr []student

	for _, s := range students {
		if f(s) {
			arr = append(arr, s)
		}
	}

	return arr
}

// Renvoie une fonction qui repond au critere de selection d'une ecole
// On a une fonction en retour
func genFilterFunc(schoolName string) func(student) bool {

	if schoolName == "A" {

		return func(s student) bool {

			// si l'etudiant a la moyenne
			if s.grade >= 5 {
				return true
			}

			// dans le cas contraire false
			return false
		}

	} else if schoolName == "B" {

		return func(s student) bool {

			// si l'etudiant a la moyenne (>= 5) et une assiduite >= 7
			if s.grade >= 5 && s.attendance >= 7 {
				return true
			}

			// dans le cas contraire false
			return false
		}

	} else {

		return func(s student) bool {
			return true
		}
	}
}

func main() {

	s1 := student{
		name:       "Bob",
		grade:      4,
		attendance: 8,
	}

	s2 := student{
		name:       "Alice",
		grade:      8,
		attendance: 4,
	}

	s := []student{s1, s2}
	f := filter(s, func(s student) bool {

		// si la note de l'etudiant est superieur a 5
		if s.grade >= 5 {
			return true
		}

		return false
	})

	fmt.Printf("Student filter simple: %v\n", f)

	// Generation d'une fonction de filtre en type de retour
	filterFunc := genFilterFunc("Z")
	f = filter(s, filterFunc)
	fmt.Printf("Students filter school Z: %v\n", f)

	// Generation d'une fonction pour l'ecole A
	filterFunc = genFilterFunc("A")
	f = filter(s, filterFunc)
	fmt.Printf("Students filter school A: %v\n", f)

	// Generation d'une fonction pour l'ecole B
	filterFunc = genFilterFunc("B")
	f = filter(s, filterFunc)
	fmt.Printf("Students filter school B: %v\n", f)
}
