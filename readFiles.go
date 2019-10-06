package main

import (
	"io/ioutil"
	"path/filepath"
)

//OperateFile struct
type OperateFile struct {
	sqlFilePath string
	filePaths   []string
}

//NewOperateFile construct
func NewOperateFile(sqlFilePath string) *OperateFile {
	o := &OperateFile{}
	o.sqlFilePath = sqlFilePath
	o.readFilePaths()
	return o
}

func (o *OperateFile) readFilePaths() {
	files, err := ioutil.ReadDir(o.sqlFilePath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}
		o.filePaths = append(o.filePaths, filepath.Join(o.sqlFilePath, file.Name()))
	}
}
