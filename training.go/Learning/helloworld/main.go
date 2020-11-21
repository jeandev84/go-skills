package main

import (
	"fmt"
	"strings"
)

// Post struct
type Post struct {
	Title     string
	Text      string
	published bool
}

// Headline return un format affichant le titre et un text de 50 characteres
func (p Post) Headline() string {
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:50]) // titre et les 50 premieres characteres
}

// Return l'etat de characteres de publication
// en lecture on est pas obliger d'utiliser une value receiver
/*
func (p *Post) Published() bool {
	return p.published
}
*/
func (p Post) Published() bool {
	return p.published
}

// Modifier l'etat de publication
func (p *Post) Publish() {
	p.published = true
}

// Remettre a l'etat non publier
func (p *Post) Unpublish() {
	p.published = false
}

// Mettre le titre en Majuscule
func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

// Entry point of Application
func main() {

	p := Post{
		Title: "Go release",
		Text: `
		Go was conceived in 2007 to improve programming productivity at Google, 
		in an era of multicore processors, computer networks, 
		and large codebases.[17] The designers wanted to resolve criticisms of other languages, while retaining many of their useful characteristics:[18]
		`,
	}

	// fmt.Println(p)
	fmt.Println(p.Headline())
	fmt.Printf("Post published? %v\n", p.Published())
	p.Publish()
	fmt.Printf("Post published? %v\n", p.Published())

	pythonPost := &Post{
		Title: "Python Intro",
		Text: `
		Go was conceived in 2007 to improve programming productivity at Google, 
		in an era of multicore processors, computer networks, 
		and large codebases.[17] The designers wanted to resolve criticisms of other languages, while retaining many of their useful characteristics:[18]
	   `,
	}

	UpperTitle(pythonPost) // or UpperTitle(&p) fmt.Println(p.Headline())
	fmt.Println(pythonPost.Headline())
}

/*

type User struct {
	Name string
}

func (u User) SayHello() {
	fmt.Printf("Hello %v!\n", u.Name)
}


func (u *User) UpdateName(name string) {
    u.Name = name
}


u := User{"Bob"}
u.SayHello() // output: Hello Bob

u.UpdateName("Alice") // u.Name == "Alice"

===================================================

Quand est-ce qu'utiliser une VALUE RECEIVER ou un POINTEUR RECEIVER ?

* Value receiver
  - est utilisee quand le coup de la copie est faible on ne veut pas pouvoir modifier les donnees)

*  Pointer receiver
  - est utilisee quand on veut modifier des donnees
	et que la consomation memoire sera plus faible

1. func (u User) SayHello() string { ... }  VALUE RECEIVER
2. func (u *User) SayHello() string { ... } POINTER RECEIVER

===================================================

type User struct {
    Name string
}
​
func (u *User) PrefixName(prefix string) {
    u.Name = fmt.Sprintf("%v %v", prefix, u.Name)
}
​
u := User{"Bob"}
u.PrefixName("Mr.")
*/
