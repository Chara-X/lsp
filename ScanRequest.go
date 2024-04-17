package lsp

import (
	"bytes"
	"strconv"
)

func ScanRequest(stream []byte, atEOF bool) (advance int, token []byte, err error) {
	if before, after, ok := bytes.Cut(stream, []byte("\r\n\r\n")); ok {
		if contentLength, _ := strconv.Atoi(string(before)[len("Content-Length: "):]); len(after) >= contentLength {
			return len(before) + contentLength + 4, after[:contentLength], nil
		}
	}
	return 0, nil, nil
}
