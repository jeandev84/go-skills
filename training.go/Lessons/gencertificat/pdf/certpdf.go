package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
	"training.go/gencertificat/cert"
)

/*
 Generateur de PDF
 OutputDir est le chemin du dossier ou sera sauvegarder nos fichiers pdf
*/
type PdfSaver struct {
	OutputDir string
}

// New PdfSaver
func New(outputdir string) (*PdfSaver, error) {

	var p *PdfSaver

	// Creer le dossier s'il n'existe pas et tous les sous dossiers
	// directory, mode de permission
	err := os.MkdirAll(outputdir, os.ModePerm)

	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: outputdir,
	}

	return p, nil
}

// Implementation de la methode Save() de l'interface
func (p *PdfSaver) Save(cert cert.Cert) error {

	// Longueur: mm (millimetre)
	// Format: A4
	// Dossier de police de caracteres : ""
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Background
	background(pdf)

	// --
	// Header
	header(pdf, &cert)

	// Saut de ligne
	pdf.Ln(30)

	// --
	// Body
	// I : Italic, fontsize: 20, C : center
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")

	// Saut de ligne
	pdf.Ln(30)

	// Body - Student name
	// Font-Familly: Times, B: Bold, 20 : Font Size
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")

	// Saut de ligne
	pdf.Ln(30)

	// Body - Participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")

	// Saut de ligne
	pdf.Ln(30)

	// Body - Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	// --
	// Footer
	footer(pdf)

	// Save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle) // nom du fichier
	path := path.Join(p.OutputDir, filename)           // chemin d'enregistrement

	err := pdf.OutputFileAndClose(path)

	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

// Make PDF Background
func background(pdf *gofpdf.Fpdf) {

	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	// Obtenir les dimenssions de la page
	pageWidth, pageHeight := pdf.GetPageSize()

	// Ajouter des options d'image de fond
	// Voir la documentation de gofpdf methode ImageOptions()
	pdf.ImageOptions("img/background.png", 0, 0, pageWidth, pageHeight, false, opts, 0, "")
}

// Make PDF Header
func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"

	// First Gopher
	pdf.ImageOptions(filename,
		x+margin, 20,
		imageWidth, 0,
		false, opts, 0, "")

	// Second Gopher
	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename,
		x-margin, 20,
		imageWidth, 0,
		false,
		opts,
		0,
		"")

	// Set Title, font,  C - center, L - left, R -right
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

// Make PDF Footer
func footer(pdf *gofpdf.Fpdf) {

	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	imageWidth := 50.0
	filename := "img/stamp.png"
	pageWidth, pageHeight := pdf.GetPageSize()

	x := pageWidth - imageWidth - 20.0
	y := pageHeight - imageWidth - 10

	pdf.ImageOptions(filename,
		x, y,
		imageWidth, 0,
		false, opts, 0, "")
}
