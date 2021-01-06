// Package tar implements access to tar archives.
// Tape archives (tar) are a file format for storing a sequence of files that
// can be read and written in a streaming manner.
package tar

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"os"
	"time"
)

// BufferFile implements a file zip compression with buffer.
type BufferFile struct {
	filename string
	buffer   *bytes.Buffer
	tw       *tar.Writer
}

// NewBufferFile returns a new BufferFile.
func NewBufferFile(filename string) *BufferFile {
	var buf bytes.Buffer
	return &BufferFile{
		filename: filename,
		buffer:   &buf,
		tw:       tar.NewWriter(&buf),
	}
}

// AddFile used to add file to tar archive, filename is compressed file
// buf is file content.
func (bf *BufferFile) AddFile(filename string, content []byte) error {
	err := bf.tw.WriteHeader(&tar.Header{
		Name:    filename,
		Mode:    0600,
		Size:    int64(len(content)),
		ModTime: time.Now(),
	})
	if err != nil {
		return err
	}
	_, err = bf.tw.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// Close finishes writing the tar file.
// It does not close the underlying writer.
func (bf *BufferFile) Close() error {
	return bf.tw.Close()
}

// Buffer returns tar archive content
func (bf *BufferFile) Buffer() *bytes.Buffer {
	return bf.buffer
}

// NoBufferFile implements a file tar compression.
// Suitable for saving compressed files locally.
type NoBufferFile struct {
	filename string
	file     *os.File
	tw       *tar.Writer
}

// NewBufferFile returns a new NoBufferFile.
func NewNoBufferFile(filename string) *NoBufferFile {
	zipFile := new(NoBufferFile)
	zipFile.filename = filename
	file, _ := os.Create(filename)
	zipFile.file = file
	zipFile.tw = tar.NewWriter(file)
	return zipFile
}

// AddFile used to add file to tar archive, filename is compressed file
// buf is file content.
func (nbf *NoBufferFile) AddFile(filename string, content []byte) error {
	err := nbf.tw.WriteHeader(&tar.Header{
		Name:    filename,
		Mode:    0600,
		Size:    int64(len(content)),
		ModTime: time.Now(),
	})
	if err != nil {
		return err
	}
	_, err = nbf.tw.Write(content)
	if err != nil {
		return err
	}
	return nil
}

// Close finishes writing the tar file.
// It does not close the underlying writer.
func (nbf *NoBufferFile) Close() error {
	if err := nbf.tw.Close(); err != nil {
		return err
	}
	return nbf.file.Close()
}

// Content returns tar archive content
func (nbf *NoBufferFile) Content() ([]byte, error) {
	return ioutil.ReadAll(nbf.file)
}
