//go:build unused
package util

import (
	"bufio"
	"io"
	"strings"
)

func isEOF(err error) bool {
	if err != nil {
		switch err {
		case io.EOF:
			return true
		default:
			panic(err)
		}
	}
	return false
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	check(err)
	return v
}

// Pipe function, to mutate stream.
func trimSpace(r io.Reader) io.Reader {
	pr, pw := io.Pipe()
	bufr := bufio.NewReader(r)

	go func() {
		for {
			s, err := bufr.ReadString('\n')
			if isEOF(err) {
				break
			}

			Must(
				pw.Write([]byte(strings.TrimSpace(s))),
			)
		}
	}()
	return pr
}
