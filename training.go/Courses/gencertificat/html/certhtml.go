package html

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"training.go/gencertificat/cert"
)

/**
  Struct HtmlSaver
  Il faudra creer une struct qui va se partager entre HtmlSaver et PdfSaver
  Car il on en commun OutputDir et les memes methodes
*/
type HtmlSaver struct {
	OutputDir string
}

// New HtmlSaver
func New(outputdir string) (*HtmlSaver, error) {

	var h *HtmlSaver

	// Creer le dossier s'il n'existe pas et tous les sous dossiers
	// directory, mode de permission
	err := os.MkdirAll(outputdir, os.ModePerm)

	if err != nil {
		return h, err
	}

	h = &HtmlSaver{
		OutputDir: outputdir,
	}

	return h, nil
}

func (p *HtmlSaver) Save(cert cert.Cert) error {

	t, err := template.New("certificate").Parse(tpl)

	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v.html", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	err = t.Execute(f, cert)

	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v'\n", path)

	return nil
}

var tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
        <title>{{.LabelTitle}}</title>
        <style>
            body {
                text-align: center;
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
            }

            h1 {
                font-size: 3em;
            }

        </style>
    </head>
	<body>
	    <img src="">
        <h1>{{.LabelCompletion}}</h1>
        <h2><em>{{.LabelPresented}}</em></h2>
        <h1>{{.Name}}</h1>
        <h2>{{.LabelParticipation}}</h2>
        <p>
            <em>{{.LabelDate}}</em>
        </p>
    </body>
</html>
`
