package rslice

import (
	"io"
	"bytes"
	"os"
	"strings"
)



type ReaderSlice interface {
	Len() int
	GetR(int) io.Reader
}

type WriterSlice interface {
	Len() int
	GetW(int) io.Writer
}



type IoReaderSlice []io.Reader

func (r IoReaderSlice) Len() int {
	return len(r)
}

func (r IoReaderSlice) GetR(i int) io.Reader {
	return r[i]
}



type IoWriterSlice []io.Writer

func (w IoWriterSlice) Len() int {
	return len(w)
}

func (w IoWriterSlice) GetW(i int) io.Writer {
	return w[i]
}


type IoReadWriterSlice []io.ReadWriter

func (w IoReadWriterSlice) Len() int {
	return len(w)
}

func (w IoReadWriterSlice) GetW(i int) io.Writer {
	return w[i]
}

func (r IoReadWriterSlice) GetR(i int) io.Reader {
	return r[i]
}


type BytesBufferSlice []*bytes.Buffer

func (w BytesBufferSlice) Len() int {
	return len(w)
}

func (w BytesBufferSlice) GetW(i int) io.Writer {
	return w[i]
}

func (r BytesBufferSlice) GetR(i int) io.Reader {
	return r[i]
}


type OsFileSlice []*os.File

func (w OsFileSlice) Len() int {
	return len(w)
}

func (w OsFileSlice) GetW(i int) io.Writer {
	return w[i]
}

func (r OsFileSlice) GetR(i int) io.Reader {
	return r[i]
}


type BytesReaderSlice []*bytes.Reader

func (w BytesReaderSlice) Len() int {
	return len(w)
}

func (r BytesReaderSlice) GetR(i int) io.Reader {
	return r[i]
}


type StringsBuilderSlice []*strings.Builder

func (w StringsBuilderSlice) Len() int {
	return len(w)
}

func (w StringsBuilderSlice) GetW(i int) io.Writer {
	return w[i]
}

