package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//GetTarGetFilePath 获取目标文件路径
func GetTarGetFilePath(currentPath string, postfix string) string {
	filePath := ""
	err := filepath.Walk(currentPath, func(path string, f os.FileInfo, err error) error {
		ok := strings.HasSuffix(path, postfix)
		if ok {
			fmt.Println("path=", path)
			filePath = path
			return errors.New("")
		}
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return filePath
}

//GetCurrentDirectory 获取当前绝对路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("GetCurrentDirectory %v\n", err)
	}
	return dir
}

//CopyFile 复制文件
func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("CopyFile " + err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("CopyFile %v" + err.Error())
		return
	}
	defer dstFile.Close()
	fmt.Println("CopyFile Success")
	return io.Copy(dstFile, srcFile)
}

//DeleteFile jacoco.exec
func DeleteFile() {
	file := "jacoco.exec"
	err := os.Remove(file)
	if err != nil {
		fmt.Println("file remove Error!")
		fmt.Printf("%s", err)
	} else {
		fmt.Print("file jacoco.exec remove OK!")
	}
}

func main() {
	DeleteFile()
	tarGetFilePath := GetTarGetFilePath(GetCurrentDirectory(), "exec")
	CopyFile(tarGetFilePath, "jacoco.exec")
}
