package task

import (
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

// ChanTask implemente la methode process de Tasker
func (c *ChanTask) Process() error {

	size := len(c.files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)

	return nil
}
