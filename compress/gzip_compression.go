package compress

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

type GzipCompression struct {
}

func (GzipCompression) Compress(input []byte) []byte {
	buffer := bytes.NewBuffer(nil)
	gzipWriter := gzip.NewWriter(buffer)

	gzipWriter.Write(input)
	gzipWriter.Close()
	compressed := buffer.Bytes()

	return compressed
}

func (GzipCompression) Decompress(input []byte) ([]byte, error) {
	gzipReader, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}

	gzipReader.Close()

	return ioutil.ReadAll(gzipReader)
}
