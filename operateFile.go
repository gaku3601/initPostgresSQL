package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//OperateFile struct
type OperateFile struct {
	sqlFilePath string
	fileInfos   []*FileInfo
}

// FileInfo struct
type FileInfo struct {
	filePath         string
	fileNameForSorts string
	fileName         string
}

//NewOperateFile construct
func NewOperateFile(sqlFilePath string) *OperateFile {
	o := &OperateFile{}
	o.sqlFilePath = sqlFilePath
	o.readFilePaths()
	o.nameingVersionName()
	o.sortMigrationFile()
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
		o.fileInfos = append(o.fileInfos, &FileInfo{filePath: filepath.Join(o.sqlFilePath, file.Name()), fileName: file.Name()})
	}
}

func (o *OperateFile) nameingVersionName() {
	for _, fileInfo := range o.fileInfos {
		rep := regexp.MustCompile(`^V(\d*?)\.(\d*?)\.(\d*?)_`)
		result := rep.FindAllStringSubmatch(fileInfo.fileName, -1)
		if result != nil {
			major := parseStrToInt(result[0][1])
			minor := parseStrToInt(result[0][2])
			build := parseStrToInt(result[0][3])
			fileInfo.fileNameForSorts = strings.Replace(fileInfo.fileName, result[0][0],
				fmt.Sprintf("V%02d.%02d.%02d_", major, minor, build), 1)
			continue
		}
		rep = regexp.MustCompile(`^V(\d*?)_(\d*?)_`)
		result = rep.FindAllStringSubmatch(fileInfo.fileName, -1)
		if result != nil {
			major := parseStrToInt(result[0][1])
			minor := parseStrToInt(result[0][2])
			fileInfo.fileNameForSorts = strings.Replace(fileInfo.fileName, result[0][0],
				fmt.Sprintf("V%02d.%02d_", major, minor), 1)
			continue
		}
	}
}

func parseStrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (o *OperateFile) sortMigrationFile() {
	sort.Slice(o.fileInfos, func(i, j int) bool {
		return o.fileInfos[i].fileNameForSorts < o.fileInfos[j].fileNameForSorts
	})
}
