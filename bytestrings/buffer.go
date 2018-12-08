package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer .
func Buffer(s string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(s))
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
