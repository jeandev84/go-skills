package main

import (
	"fmt"
	"io/ioutil"
)

/*
Gestion d'erreur

func MyFunc() (int, error) { ... }

v, err := MyFunc()

// nil signifit si la variable n'a pas ete du tout definit
if err != nil {
   fmt.Printf("Error is MyFunc: %v", err)
}

v, err := MyFunc1()
if err != nil {
	return err
}


v, err := MyFunc2()
if err != nil {
	return err
}

v, err := MyFunc3()
if err != nil {
	return err
}

v, err := MyFunc4()
if err != nil {
	return err
}

============================================
Code complexe :

func MyFunc(condition bool) (int, err) {
	if (condition) {
		if (!condition2) {
			return 0, errors.New("Error 2!")
		}

		// code
		return 42, nil
	}
	return 0, errors.New("Error!")
}

Code non-early return

func MyFunc(condition bool) (int, err) {
	if (!condition) {
	   return 0, errors.New("Error!")
	}

	if (!condition) {
	   return 0, errors.New("Error 2!")
	}

	// code
	return 42, nil
}
*/

func readFile(filename string) (string, error) {

	dat, err := ioutil.ReadFile(filename)

	/*
	 Retour d'ereurs provenant de lecture de fichier
	 chaine vide "" car on a pas recu a lire le contenu d'une valeur
	*/
	if err != nil {
		return "", err
	}

	/*
	 errors.New() - permet de creer une nouvelle erreure
	 Verifier si la chaine de caracteres est vide
	*/
	if len(dat) == 0 {
		/*
		 Afficher erreur simple
		 return "", errors.New("Empty content")
		 ========================================
		 Formater une erreure
		*/
		return "", fmt.Errorf("Empty content (filename=%v)", filename)
	}

	return string(dat), nil
}

// Entry point of program
func main() {

	dat, err := readFile("text.txt")

	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return // on sort de la fonction
	}

	// si tout va bien
	fmt.Println("File content:")
	fmt.Println(dat)
}
