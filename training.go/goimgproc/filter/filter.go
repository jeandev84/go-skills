package filter

import (
	"github.com/disintegration/imaging"
)

// Filter interface
type Filter interface {
	Process(srcPath, dstPath string) error
}

// Struct Grayscale, on pouvait le nomme GrayscaleFilter
type Grayscale struct {
}

// Implementation de la methode Process de l'interface Filter
func (g Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open
}
