package main

import "fmt"

type student struct {
	name       string
	grade      int
	attendance int
}

func filter(students []student, f func(student) bool) []student {
	var arr []student
	for _, s := range students {
		if f(s) {
			arr = append(arr, s)
		}
	}
	return arr
}

func genFilterFunc(schoolName string) func(student) bool {
	if schoolName == "A" {
		return func(s student) bool {
			if s.grade >= 5 {
				return true
			}
			return false
		}
	} else if schoolName == "B" {
		return func(s student) bool {
			if s.grade >= 5 && s.attendance >= 7 {
				return true
			}
			return false
		}
	} else {
		return func(s student) bool { return true }
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

	// Les étudiants sont filtrés suivant une fonction lambda
	f := filter(s, func(s student) bool {
		if s.grade >= 5 {
			return true
		}
		return false
	})

	fmt.Printf("Students Filter simple: %v\n", f)

	// Génération d'une fonction en type de retour
	filterFunc := genFilterFunc("A")
	f = filter(s, filterFunc)
	fmt.Printf("Students Filter school A: %v\n", f)

	filterFunc = genFilterFunc("B")
	f = filter(s, filterFunc)
	fmt.Printf("Students Filter school B: %v\n", f)
}
