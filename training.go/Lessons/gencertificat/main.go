package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/gencertificat/cert"
	"training.go/gencertificat/html"
	"training.go/gencertificat/pdf"
)

// Entry point of program
// $ go build (will be create file gencertificat)
// $ go run main.go
// GENERATE FILE PDF: $ ./gencertificat -type pdf
// GENERATE FILE PDF: $ ./gencertificat -type html
// CREATE file students.csv : $ touch students.csv
/*
# PDF
~/go/src/training.go/gencertificat$ ./gencertificat -type pdf -file students.csv
Saved certificate to 'output/GOLANG PROGRAMMING COURSE Certificate - BOB DYLAN.pdf'
Saved certificate to 'output/GOLANG NETWORK COURSE Certificate - ALICE LIDDELL.pdf'
Saved certificate to 'output/GOLANG WEB COURSE Certificate - BOBETTE LA BROSSE.pdf'
Saved certificate to 'output/GOLANG IMAGE COURSE Certificate - JOHN DOE.pdf'


# HTML
~/go/src/training.go/gencertificat$ ./gencertificat -type html -file students.csv
Saved certificate to 'output/GOLANG PROGRAMMING COURSE Certificate - BOB DYLAN.html'
Saved certificate to 'output/GOLANG NETWORK COURSE Certificate - ALICE LIDDELL.html'
Saved certificate to 'output/GOLANG WEB COURSE Certificate - BOBETTE LA BROSSE.html'
Saved certificate to 'output/GOLANG IMAGE COURSE Certificate - JOHN DOE.html'
*/
func main() {

	// Parse file
	file := flag.String("file", "", "CSV file input")

	// CLI
	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	flag.Parse()

	// tester si le fichier est vide
	if len(*file) <= 0 {
		fmt.Printf("Invalid file. got=%v\n", *file)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error

	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output type. got=%v\n", *outputType)
	}

	/*
	 saver, err = pdf.New("output")
	 saver, err = html.New("output")
	*/

	if err != nil {
		fmt.Printf("Could not create generator: %v", err)
		os.Exit(1)
	}

	/*
		Creer un certificat
		c, err := cert.New("Golang programming", "Bob Dylan", "2018-06-21")

		if err != nil {
			fmt.Printf("Error during certificate creation: %v", err)
			os.Exit(1)
		}

		saver.Save(*c)
	*/

	certs, err := cert.ParseCSV(*file)

	if err != nil {
		fmt.Printf("Could not parse CSV file : %v", err)
		os.Exit(1)
	}

	for _, c := range certs {

		err = saver.Save(*c)

		if err != nil {
			fmt.Printf("Could not save Cert. got=%v\n", err)
		}
	}
}
