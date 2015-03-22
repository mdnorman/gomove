package gomove

import "fmt"

func MoveDirectory(srcDir string, destDir string) {
	fmt.Printf("Moving directory '%s' to '%s'", srcDir, destDir)
}

func MoveFile(srcFile string, destDir string) {
	fmt.Println("Moving file '%s' to '%s'", srcFile, destDir)
}
