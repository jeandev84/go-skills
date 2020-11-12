# PROJECT GENERATOR PDF and HTML
 - But du projet est d' ecrire un programme qui permet de generer des certificats d'accomplissements


# READ CSV
``$ cat students.csv``

# GENERER UN FICHIER PAR TYPE
``$ ./gencert -type html -file students.csv``

# 3 used packages (cert, pdf, html)
- cert
``Logique metier``
``Tests unitaires``
``Parsing CSV``


- pdf
`` Bibliotheque de generation de PDF jung-kurt/gofpdf``
``Integration et placement d'images``


- html
`` Moteur de template Go``
``Generation d'un fichier HTML Simple``

# GO TIME PACKAGE 
``https://golang.org/pkg/time/#pkg-examples``


# PACKAGE GOFPDF for generate pdf
- Site officiel" https://github.com/jung-kurt/gofpdf/
``go get github.com/jung-kurt/gofpdf``

# PACKAGE go html template 
- Site officiel " https://golang.org/pkg/html/template/ 


# FORMAT CSV (Comma Separate Value)


# ALLER PLUS LOIN
* Idees:
- Supporter un autre format d'entree (JSON, XML) comme le CSV
   -- (Utiliser une interface GO pour abstraire le tout)

- Ameliorer le rendu HTML (image CSS)

- Packager le resultat HTML dans un fichier zip 
   -- (Utiliser le package archive/zip)

- Supporter un PDF en mode portrait