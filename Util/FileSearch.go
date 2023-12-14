package Util

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func readAllFilesInFolder(folderPath string) ([]string, error) {
	var fileContents []string

	// 列出文件夹中的所有文件
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	// 逐一读取每个文件的内容
	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())

		// 读取文件内容
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		// 将文件内容添加到数组中
		fileContents = append(fileContents, string(content))
	}

	return fileContents, nil
}
func searchFilesByKeyword(folderPath, keyword string) ([]string, error) {
	allContents, err := readAllFilesInFolder(folderPath)
	if err != nil {
		return nil, err
	}

	var matchingContents []string

	// 在所有文件内容中搜索包含关键字的文件
	for _, content := range allContents {
		if strings.Contains(content, keyword) {
			matchingContents = append(matchingContents, content)
		}
	}

	return matchingContents, nil
}

func GetFiles(folder string) []fs.FileInfo {
	files, _ := ioutil.ReadDir(folder)
	return files
}
