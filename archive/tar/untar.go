package tar

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
)

// UnCompressedFile used to save unzipped file.
type UnCompressedFile struct {
	Filename string // 文件名称
	Content  []byte // 文件内容
}

// Unzip will open the tar file specified by name and
// return compressed file content.
func Unzip(filename string) ([]UnCompressedFile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	tr := tar.NewReader(file)
	var ucFiles = make([]UnCompressedFile, 0)
	for {
		var ucFile UnCompressedFile
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		ucFile.Filename = header.Name

		content, err := ioutil.ReadAll(tr)
		if err != nil {
			return nil, err
		}
		ucFile.Content = content
		ucFiles = append(ucFiles, ucFile)
	}
	return ucFiles, nil
}
