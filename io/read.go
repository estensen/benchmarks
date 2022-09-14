package io

import (
	"bytes"
	"io"
)

func IOReadAll(reader io.Reader) []byte {
	content, _ := io.ReadAll(reader)
	return content
}

func IOCopy(reader io.Reader) []byte {
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, reader)
	return buf.Bytes()
}

func IOBufferReadFrom(reader io.Reader) []byte {
	var buf bytes.Buffer
	buf.ReadFrom(reader)
	return buf.Bytes()
}
