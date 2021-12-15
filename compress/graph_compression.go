package compress

import "encoding/base64"

type GraphCompression struct {
}

func (GraphCompression) CompressGraphData(input []byte) string {
	compressed := GzipCompression{}.Compress(input)
	return base64.StdEncoding.EncodeToString(compressed)
}

func (GraphCompression) DecompressGraphData(input string) ([]byte, error) {
	compressed, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return GzipCompression{}.Decompress(compressed)
}
