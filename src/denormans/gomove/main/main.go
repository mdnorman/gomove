package main

import (
	"denormans/gomove"
	"flag"
	"fmt"
	"os"
	"path"
)

type ErrorCode int

const (
	UsageError ErrorCode = 1 << iota
)

func main() {
	var srcDir string
	var destParentDir string

	flag.StringVar(&srcDir, "from", "", "The source directory to move")
	flag.StringVar(&destParentDir, "to", "", "The destination parent directory for the directory being moved")

	flag.Parse()

	if len(srcDir) == 0 {
		fmt.Fprintln(os.Stderr, "The source directory is required")
		flag.PrintDefaults()
		os.Exit(int(UsageError))
	}

	srcDirInfo, err := os.Stat(srcDir)
	if err != nil {
		ExitWithError(err, UsageError, "Couldn't get info on the source directory to move:", srcDirInfo)
	}

	if !srcDirInfo.IsDir() {
		ExitWithError(nil, UsageError, "Source is not a directory:", srcDir)
	}

	if len(destParentDir) == 0 {
		fmt.Fprintln(os.Stderr, "The destination parent directory is required")
		flag.PrintDefaults()
		os.Exit(int(UsageError))
	}

	destParentDirInfo, err := os.Stat(destParentDir)
	if err != nil {
		err = os.MkdirAll(destParentDir, srcDirInfo.Mode())
		if err != nil {
			ExitWithError(err, UsageError, "Couldn't get info on the destination parent directory to move to:", destParentDirInfo)
		}

		destParentDirInfo, err = os.Stat(destParentDir)
	}

	if !destParentDirInfo.IsDir() {
		ExitWithError(nil, UsageError, "Destination is not a directory:", destParentDir)
	}

	destDir := path.Join(destParentDir, path.Base(srcDir))

	gomove.MoveDirectory(srcDir, destDir)
}

func ExitWithError(err error, exitCode ErrorCode, message ...interface{}) {
	fmt.Fprintln(os.Stderr, message...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(int(exitCode))
}
