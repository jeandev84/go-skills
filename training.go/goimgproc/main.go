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
/*
$ go run main.go -src imgs/ -dst output -filter blur -task waitgrp
*/
func main() {

	var srcDir = flag.String("src", "", "Input Directory")
	var dstDir = flag.String("dst", "", "Output Directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	var taskType = flag.String("task", "waitgrp", "waitgrp/channel")
	var poolSize = flag.Int("poolsize", 4, "Workers pool size for the channel task")

	flag.Parse()

	var f filter.Filter = filter.Grayscale{}

	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
		/*default:
		fmt.Println("Filter not found")*/
	}

	/*
		t := task.NewWaitGrpTask("./imgs", "output", f)
		t := task.NewChanTask("./imgs", "output", f, 4)
		t := task.NewChanTask(*srcDir, *dstDir, f, 16)
	*/

	var t task.Tasker

	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	}

	start := time.Now() // debut du temps du process
	t.Process()
	elapsed := time.Since(start) // temps ecouler

	fmt.Printf("Image processing took %s\n", elapsed)
}
