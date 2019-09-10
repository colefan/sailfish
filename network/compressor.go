package network

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// CompressorType 压缩类型
type CompressorType byte

const (
	// None 无压缩
	None CompressorType = iota
	// Gzip gzip压缩
	Gzip
)

// Compressor 压缩接口
type Compressor interface {
	Zip([]byte) ([]byte, error)
	Unzip([]byte) ([]byte, error)
}

// GzipCompressor gzip 压缩算法
type GzipCompressor struct {
}

// Zip zip the data
func (c *GzipCompressor) Zip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unzip unzip the data
func (c *GzipCompressor) Unzip(data []byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return data, err
}
