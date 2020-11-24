package main

import (
	"flag"
	"fmt"
	"time"

	"training.go/goimgproc/filter"
	"training.go/goimgproc/task"
)

// Entry point
// $ go run main.go
func main() {

	var srcDir = flag.String("src", "", "Input Directory")
	var dstDir = flag.String("dst", "", "Output Directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	flag.Parse()

	var f filter.Filter = filter.Grayscale{}

	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	/*
		t := task.NewWaitGrpTask("./imgs", "output", f)
		t := task.NewChanTask("./imgs", "output", f, 4)
	*/

	t := task.NewChanTask(*srcDir, *dstDir, f, 16)

	start := time.Now() // debut du temps du process
	t.Process()
	elapsed := time.Since(start) // temps ecouler

	fmt.Printf("Image processing took %s\n", elapsed)
}
