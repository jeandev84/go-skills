package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine searches for old in line to replace it by new.
// It returns found=true, if the pattern was found, res with the resulting string
// and occ with the number of occurence of old
func ProcessLine(line, old, new string) (found bool, res string, occ int) {

	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)

	// initialize res a line au cas ou on n'a pas trouve de ligne
	res = line

	// si on trouve une line dans old ou en minuscule, on passe found a true
	// on augmente le nombre d'occurence par rapport au nombre d'occurence on aura a l'interieur du compte de notre lines
	// res on remplace a l'interieur de line, old et new
	// -1 nombre de limite de remplacement (-1 pour indiquer un nombre illimite)
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}

	return found, res, occ
}

// FindReplaceFile ()
// src - nom du fichier
// dst - fichier de destination
func FindReplaceFile(src string, dst string, old string, new string) (occ int, lines []int, err error) {

	// Ouvrir le fichier et verifier qu'il n'y as pas de problemes
	srcFile, err := os.Open(src)

	// traiter l'erreure imediatement
	if err != nil {
		// on utilise les variables qui on ete declare avant
		return occ, lines, err
	}

	// on ferme le fichier
	defer srcFile.Close()

	// Create le fichier de destination
	dstFile, err := os.Create(dst)

	if err != nil {
		// on utilise les variables qui on ete declare avant
		return occ, lines, err
	}

	defer dstFile.Close()

	// Initialize old
	old = old + " "
	new = new + " "

	// Index de Line
	lineIdx := 1

	// On va lire le contenu du fichier en initialisant un scanner de lecture de fichier
	scanner := bufio.NewScanner(srcFile)

	// On va initialiser un writer qui ecrire ligne par ligne le contenu du fichier
	writer := bufio.NewWriter(dstFile)

	// On vide le buffer
	defer writer.Flush()

	// Parcourt du scanner
	for scanner.Scan() {

		// scanner.Text() - text en une ligne donnee
		// old - l'ancien mot
		// new - nouveau mot
		found, res, o := ProcessLine(scanner.Text(), old, new)

		// si on a trouve le mot
		if found {
			// on va augmente l'occurence du compter global par rapport a celui qui nous a ete renvoye
			occ += o

			// on ajoute a lines l'index de la ligne
			lines = append(lines, lineIdx)
		}

		// fmt.Println(res)
		fmt.Fprintf(writer, res)
		lineIdx++
	}

	// si tout se passe bien on retourne nil pour dire qu'on n'a pas eu d'erreurs
	return occ, lines, nil
}

// Entry point of program
func main() {

	// Call function ProcessLine()
	/*
		found, res, occ := ProcessLine(
			"Go was conceived in 2007 to improve programming productivity at Google, in an era of multicore processors, computer networks, and large codebases.[17] The designers wanted to resolve criticisms of other languages, while retaining many of their useful characteristics:[18]",
			"Go",
			"Python",
		)

		fmt.Println(found, res, occ)
	*/

	// Call function FindReplaceLine()
	old := "Go"
	new := "Python"

	occ, lines, err := FindReplaceFile("wikigo.txt", "wikipyton.txt", old, new)

	// traiter l'erreur en premier
	if err != nil {
		fmt.Printf("Error while executing find replace: %v\n", err)
	}

	// len() determine la taille d'un tableau
	fmt.Println("== Summary ==")
	defer fmt.Println("== End of Summary ==") // Ceci s'executera a pres avoir executer tout le programme

	fmt.Printf("Number of occurences of %v: %v\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))

	// formattage des lines:
	fmt.Print("Lines: [ ")
	len := len(lines)

	// on parcourt les lines
	for i, l := range lines {
		fmt.Printf("%v", l)

		// si on est inferieur a l'index (len -1)
		if i < len-1 {

			// on ajoute mon separateur
			fmt.Printf(" - ")
		}
	}

	fmt.Println(" ]")
}

/*
Function 1:
func FindReplaceFile(src, old, new string) (occ int, lines []int, err error)
src : nom du fichier source
old/new: ancien mot / nouveau mot
occ: nombre d'occurences de old
lines: slice des numeros de lignes ou old a ete trouve
err: indique erreur de la fonction

Function 2:
func ProcessLine(line, old, new string) (found bool, res string, occ int)
line: ligne a traiter
old/new: ancien mot / nouveau mot
found: vrai si moins occurence traite
res: resultat de remplacement (res == line si aucun changement!)
occ: nombre d'occurence de old dans la ligne

Outils - Package bufio
Lire de fichier ligne par ligne

scanner := bufio.NewScanner(srcFile)

for scanner.Scan () {
	t := scanner.Text() // t contient une ligne
}

Outils - package strings
Manipuler des strigns

c:= strings.Contains("go ruby java", "go") c == true
cnt := strings.Count("go go go", "go") cnt == 3
res := strings.Replace("go", "python", 0) res == python

===========================================================
Contains() - permet de savoir si un text contient une chaine de characters ou non)
Count()     - permet de savoir le nombres d'occurences ou non
Replace()   - permet de remplace la chaine de characteres

===========================================================
Bonus
 * Remplacer le en minuscule
   - Go et go ==> Python et python
 * Stocker le resultat dans un fichier
   - Modifier legerement FindReplaceFile()

============================================================
Outils - package bufio
Code a utiliser pour ecrire ligne par ligne dans un fichier

dstFile (destination file)
writer := bufio.NewWriter(dstFile)
defer writer.Flush()
permet de rendre l'ecriture du disque definitif et de vider le buffer contenu dans writer
fmt.Fprintln(writer, "Texte d'une ligne")

-------------------------------------------------------------
line - ligne  traiter
old  - l'ancien mot
new  - nouveau mot
---------------------------
found - pour savoir si on a trouve le mot ou pas
res   - le resultat de la chaine modifier
occ   - le nombre d'occurence du mot
---------------------------
Utiliser les noms courts en go

===================================================================================
Resultat apres execution:

$ go run main.go
Python (programming language)
From Wikipedia, the free encyclopedia
  (Redirected from Golang)
Jump to navigationJump to search
For the language released in 2003 by McCabe and Clark, see Go! (programming language).
"Google Go" redirects here. For the computer program by Google to play the board game Go, see AlphaGo.
￼
Python was conceived in 2007 to improve programming productivity at Google, in an era of multicore processors, computer networks, and large codebases.[17] The designers wanted to resolve criticisms of other languages, while retaining many of their useful characteristics:[18]

Static typing and efficiency (like C++ or Java)
Productivity and ease of use (like Python or JavaScript)[19]
High-performance networking and multiprocessing
The designers cited their shared dislike of C++ as a primary motivation for designing a new language.[20][21][22]

Python was publicly announced in November 2009,[23] and version 1.0 was released in March 2012.[24][25] Python is widely used in production at Google[26] and in many other organizations and open-source projects.

In April 2018, the original lopython (Gopher mascot) was replaced with a stylized GO slanting right with trailing streamlines. However, the mascot remained the same.[27]

In August 2018, the Python principal contributors published two ″draft designs″ for new language features, Generics and Error Handling, and asked Python users to submit feedback on them.[28][29] Lack of support for generic programming and the verbosity of error handling in Python 1.x had drawn considerable criticism.

￼
Language design[edit]
Python is recognizably in the tradition of C, but makes many changes to improve brevity, simplicity, and safety. The language consists of:

A syntax and environment adopting patterns more common in dynamic languages:[40]
Optional concise variable declaration and initialization through type inference (x := 0 not int x = 0; or var x = 0;).
Fast compilation times.[41]
Remote package management (python get)[42] and online package documentation.[43]
Distinctive approaches to particular problems:
Built-in concurrency primitives: light-weight processes (goroutines), channels, and the select statement.
An interface system in place of virtual inheritance, and type embedding instead of non-virtual inheritance.
A toolchain that, by default, produces statically linked native binaries without external dependencies.
A desire to keep the language specification simple enough to hold in a programmer's head,[44] in part by omitting features which are common in similar languages.
== Summary ==
Number of occurences of Go: 10
Number of lines: 7
== End of Summary ==
*/
