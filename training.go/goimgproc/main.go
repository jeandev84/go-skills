package main

import (
	"fmt"
	"time"

	"training.go/goimgproc/filter"
	"training.go/goimgproc/task"
)

// Entry point
// $ go run main.go
func main() {

	/*
		f := task.BuildFileListSecond("./imgs")
		fmt.Println(f)
	*/

	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs", "output", f)

	start := time.Now() // debut du temps du process
	t.Process()
	elapsed := time.Since(start) // temps ecouler

	fmt.Printf("Image processing took %s\n", elapsed)
}
