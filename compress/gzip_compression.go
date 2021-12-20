package compress

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

type GzipCompression struct {
}

func (GzipCompression) Compress(input []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	gzipWriter := gzip.NewWriter(buffer)

	_, err := gzipWriter.Write(input)
	if err != nil {
		return nil, err
	}

	gzipWriter.Close()
	compressed := buffer.Bytes()

	return compressed, nil
}

func (GzipCompression) Decompress(input []byte) ([]byte, error) {
	gzipReader, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}

	gzipReader.Close()

	return ioutil.ReadAll(gzipReader)
}
