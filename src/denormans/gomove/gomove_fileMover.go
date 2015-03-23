package gomove

import (
	"log"
	"sync"
)

type MoveFileInfo struct {
	SrcFile string
	DestDir string
	Error   error
}

type FileMover struct {
	waitGroup sync.WaitGroup

	resultChan chan MoveFileInfo
	doneChan   chan bool
	errors     []error
}

func NewFileMover() *FileMover {
	fm := &FileMover{
		resultChan: make(chan MoveFileInfo),
		doneChan:   make(chan bool),
		errors:     make([]error, 0),
	}

	fm.start()
	return fm
}

func (fm *FileMover) start() {
	go func() {
		for {
			select {
			case result := <-fm.resultChan:
				if result.Error != nil {
					fm.errors = append(fm.errors, result.Error)
				}
				fm.waitGroup.Done()

			case <-fm.doneChan:
				// all done
				return

			default:
				// nothing
			}
		}
	}()

	go func() {
		fm.waitGroup.Wait()
		fm.doneChan <- true
	}()
}

func (fm *FileMover) ProcessFile(srcFile string, destFile string) {
	fm.waitGroup.Add(1)
	go fm.processFile(MoveFileInfo{srcFile, destFile, nil})
}

func (fm *FileMover) processFile(moveFileInfo MoveFileInfo) {
	moveFileInfo.Error = MoveFile(moveFileInfo.SrcFile, moveFileInfo.DestDir)
	if moveFileInfo.Error != nil {
		log.Printf("Error moving file '%s' to '%s': %s", moveFileInfo.SrcFile, moveFileInfo.DestDir, moveFileInfo.Error)
	}

	fm.resultChan <- moveFileInfo
}

func (fm *FileMover) GetErrors() []error {
	fm.waitGroup.Wait()
	return fm.errors
}
