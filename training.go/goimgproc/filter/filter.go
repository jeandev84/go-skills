package filter

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

// Filter interface
type Filter interface {
	Process(srcPath, dstPath string) error
}

// Struct Grayscale, on pouvait le nomme GrayscaleFilter
type Grayscale struct {
}

// Struct Blur, on pouvait le nomme BlurFilter
type Blur struct {
}

// Implementation de la methode Process de l'interface Filter
func (g Grayscale) Process(srcPath, dstPath string) error {

	src, err := imaging.Open(srcPath)

	if err != nil {
		return err
	}

	dst := imaging.Grayscale(src)

	dstFile, err := os.Create(dstPath)

	if err != nil {
		return err
	}

	defer dstFile.Close()

	opts := jpeg.Options{
		Quality: 90,
	}

	return jpeg.Encode(dstFile, dst, &opts)
}

// Implementation de la methode Process de l'interface Filter
func (b Blur) Process(srcPath, dstPath string) error {

	src, err := imaging.Open(srcPath)

	if err != nil {
		return err
	}

	dst := imaging.Blur(src, 3.5)

	dstFile, err := os.Create(dstPath)

	if err != nil {
		return err
	}

	defer dstFile.Close()

	opts := jpeg.Options{
		Quality: 90,
	}

	return jpeg.Encode(dstFile, dst, &opts)
}
