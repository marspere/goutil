// Package zip provides support for reading and writing ZIP archives.
package zip

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"time"
)

// BufferFile implements a file zip compression with buffer.
// It is suitable for not creating files locally,
// but for file cloud storage through buffer.
type BufferFile struct {
	filename string
	buf      *bytes.Buffer
	zw       *zip.Writer
}

// NewBufferFile returns a new BufferFile.
func NewBufferFile(filename string) *BufferFile {
	zipFile := new(BufferFile)
	zipFile.filename = filename
	zipFile.buf = new(bytes.Buffer)
	zipFile.zw = zip.NewWriter(zipFile.buf)
	return zipFile
}

// AddFile used to add file to zip archive, filename is compressed file
// buf is file content.
func (bf *BufferFile) AddFile(filename string, content []byte) error {
	header := &zip.FileHeader{
		Name:     filename,
		Method:   zip.Deflate,
		Modified: time.Now(),
	}
	zipWriter, err := bf.zw.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = zipWriter.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// Close finishes writing the zip file.
// It does not close the underlying writer.
func (bf *BufferFile) Close() error {
	return bf.zw.Close()
}

// Buffer returns zip archive content
func (bf *BufferFile) Buffer() *bytes.Buffer {
	return bf.buf
}

// NoBufferFile implements a file zip compression.
// Suitable for saving compressed files locally.
type NoBufferFile struct {
	filename string
	file     *os.File
	zw       *zip.Writer
}

// NewBufferFile returns a new NoBufferFile.
func NewNoBufferFile(filename string) *NoBufferFile {
	zipFile := new(NoBufferFile)
	zipFile.filename = filename
	file, _ := os.Create(filename)
	zipFile.file = file
	zipFile.zw = zip.NewWriter(file)
	return zipFile
}

// AddFile used to add file to zip archive, filename is compressed file
// buf is file content.
func (nbf *NoBufferFile) AddFile(filename string, content []byte) error {
	header := &zip.FileHeader{
		Name:     filename,
		Method:   zip.Deflate,
		Modified: time.Now(),
	}
	zipWriter, err := nbf.zw.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = zipWriter.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// Close finishes writing the zip file.
// It does not close the underlying writer.
func (nbf *NoBufferFile) Close() error {
	if err := nbf.zw.Close(); err != nil {
		return err
	}
	return nbf.file.Close()
}

// Content returns zip archive content
func (nbf *NoBufferFile) Content() ([]byte, error) {
	return ioutil.ReadAll(nbf.file)
}
