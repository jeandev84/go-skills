package cert

import (
	"fmt"
	"strings"
	"time"
)

// Constante
var MaxLenCourse = 20
var MaxLenName = 30

/*
  Course (Cours suivi)
  Name   (Nom de l'etudiant)
  Date   (Date a laquelle)

  LabelTitle         (Libelle de Title)
  LabelCompletion    (Libelle de Completion)
  LabelPresented     (Libelle de Present)
  LabelParticipation (Libelle de Participation)
  LabelDate          (Libelle de Date)
*/
type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

/*
Interface
Methode Save() prendra en parametre "Cert"
pour savoir le type de certificat a sauvegarder
*/
type Saver interface {
	Save(c Cert) error
}

/*
Function permettant de retourner un nouveau certificat
course : le cours
name   : le nom
date   : la date
*/
func New(course, name, date string) (*Cert, error) {

	// Validation de course
	c, err := validateCourse(course)

	if err != nil {
		return nil, err
	}

	// Validation de name
	n, err := validateName(name)

	if err != nil {
		return nil, err
	}

	// Formatter la date
	d, err := parseDate(date)

	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPresented:     "This Certificate is Presented To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

/*
 Fonction de validation de course
*/
func validateCourse(course string) (string, error) {

	c, err := validateStr(course, MaxLenCourse)

	if err != nil {
		return "", err
	}

	// s'il n'y as pas de suffix dans ce cas on va l'ajouter
	if !strings.HasSuffix(c, " course") {
		// concatenation
		c = c + " course"
	}

	return strings.ToTitle(c), nil
}

/*
  Validate Name
*/
func validateName(name string) (string, error) {

	n, err := validateStr(name, MaxLenName)

	if err != nil {
		return "", err
	}

	return strings.ToTitle(n), nil
}

/*
 Parse Date
*/
func parseDate(date string) (time.Time, error) {

	t, err := time.Parse("2006-01-02", date)

	if err != nil {
		return t, err
	}

	return t, nil
}

/*
  Fonction de validation de chaine de caracteres (string)
*/
func validateStr(str string, maxLen int) (string, error) {

	// trim les espaces
	c := strings.TrimSpace(str)

	// si on rentre dans cette condition on affiche une erreur
	if len(c) <= 0 {
		return c, fmt.Errorf("Invalid string. got='%s', len=%d", c, len(c))

	} else if len(c) > maxLen {
		return c, fmt.Errorf("Invalid string. got='%s', len=%d", c, len(c))
	}

	// si non, tout va bien
	return c, nil
}
