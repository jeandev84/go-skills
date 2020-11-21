/*
Pointeur
- C'est un mecanisme royal pour partager et modifier la memoire
- Un pointeur est une variable qui stock l'addresse d'une autre variable
- Les Pointeurs En Go, nous permettent de modifier les parametres des fonctions
- C'est exactement ce qu'on aimerait faire avec la methodes

============================
Memoire
   x     s    p     i
  -42    Bob  @1   -42
   @1    @2   @3    @4
=============================

x := -42
s := "Bob"
p := &x    // Creation d'un pointeur vers la variable x

(On recupere le contenu de ce que porte p, c'est a dire le contenu de @1)
i := *p    // Dereferencement de p pour recuperer la valeur de x

*/
