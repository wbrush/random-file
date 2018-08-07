package os_services

import (
	"io/ioutil"
)

func ReadDirectory(path string) (fileList []string, err error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return
	}

	for _, f := range files {
		if (!f.IsDir()) {
			fileList = append(fileList, f.Name())
		}
	}

	return
}