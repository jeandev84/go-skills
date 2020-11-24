package task

import (
	"fmt"
	"path"
	"path/filepath"

	"training.go/goimgproc/filter"
)

// Channel Task
type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {

	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
		PoolSize: poolSize,
	}
}

// job Req
type jobReq struct {
	src string
	dst string
}

// worker job
func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, results chan<- string) {

	for j := range jobs {
		fmt.Printf("worker %d, started job %v\n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		fmt.Printf("worker %d, finished job %v\n", id, j)
		results <- j.dst
	}
}

// ChanTask implemente la methode process de Tasker
func (c *ChanTask) Process() error {

	// la taile fe fichier
	size := len(c.files)

	// variable ou sera poste le channel des jobs
	jobs := make(chan jobReq, size)

	// variable ou sera poste le channel des resultats
	results := make(chan string, size)

	// on cree des workers en fonction du PoolSize
	// si PoolSize vaut 4 alors il sera generer 4 go routines
	for w := 1; w <= c.PoolSize; w++ {
		go worker(w, c, jobs, results)
	}

	// start jobs
	for _, f := range c.files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, filename)

		// Ecrire dans le channel
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}

	// Fermet le channel jobs
	close(jobs)

	// une facon de boucler sur tous les fichiers
	for range c.files {
		fmt.Println(<-results)
	}

	return nil
}
