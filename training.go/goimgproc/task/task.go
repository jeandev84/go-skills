package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Interface permettant de faire des traitements sur un fichier
type Tasker interface {
	Process() error
}

// Directory context
type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

// Method builder file list
func buildFileList(srcDir string) []string {

	files := []string{}
	fmt.Println("Generating file list...")

	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {

		// filter pour recuperer seulement les fichiers JPEG (jpg)
		if info.IsDir() || !strings.HasSuffix(path, ".jpg") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files
}

func BuildFileListSecond(srcDir string) []string {

	files := []string{}
	fmt.Println("Generating file list...")

	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {

		// filter pour recuperer seulement les fichiers JPEG (jpg)
		if info.IsDir() || !strings.HasSuffix(path, ".jpg") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files
}
