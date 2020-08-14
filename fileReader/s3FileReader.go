package fileReader

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type FileInfo struct {
	FileBytes *bytes.Reader
	FileSize  int64
	FileName  string
}

func (f *FileInfo) ReadFile(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	name := strings.Split(fname, "/")
	fmt.Println("file name is : ", name[len(name)-1])
	f.FileName = name[len(name)-1]
	f.FileSize = size
	f.FileBytes = bytes.NewReader(buffer)

	return nil
}
