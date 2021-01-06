// Package zip provides support for reading and writing ZIP archives.
package zip

import (
	"archive/zip"
	"io/ioutil"
)

// UnCompressedFile used to save unzipped file.
type UnCompressedFile struct {
	Filename string // 文件名称
	Content  []byte // 文件内容
}

// Unzip will open the zip file specified by name and
// return compressed file content.
func Unzip(filename string) ([]UnCompressedFile, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var ucFiles = make([]UnCompressedFile, 0)
	for _, file := range r.File {
		var ucFile UnCompressedFile
		ucFile.Filename = file.Name
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		content, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		ucFile.Content = content
		rc.Close()
		ucFiles = append(ucFiles, ucFile)
	}
	return ucFiles, nil
}
