package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {

	c, err := New("Golang", "Bob", "2018-05-31")

	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}

	if c == nil {
		t.Errorf("Cert should be valid reference. got=nil")
	}

	/*
		if c.Course != "Golang" {
			t.Errorf("Course name is not valid. expected='Golang', got=%v", c.Course)
		}
	*/

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. expected='GOLANG COURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {

	_, err := New("", "Bob", "2018-05-31")

	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}

}

func TestCourseTooLong(t *testing.T) {

	// Tester si la chaine de caracteres est trop longue
	course := "azertyuiopqsdfghjklmwxcvbnazertyuiosdfghjkertyuisdfghj"

	_, err := New(course, "Bob", "2018-05-31")

	if err == nil {
		t.Errorf("Error should be returned on an a too long course name (course=%s)", course)
	}

}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2018-05-31")

	if err == nil {
		t.Error("Error should be returned on an empty name")
	}
}

func TestNameTooLong(t *testing.T) {

	// Tester si la chaine de caracteres est trop longue
	name := "azertyuiopqsdfghjklmwxcvbnazertyuiosdfghjkertyuisdfghj"

	_, err := New("Golang", name, "2018-05-31")

	if err == nil {
		t.Errorf("Error should be returned on an a too long course name (name=%s)", name)
	}

}
