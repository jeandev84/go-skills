package main

import (
	"bufio"
	"os"
	"testing"
)

const FILE_PATH = "test.txt"
const LINES_COUNT = 100000
const TEXT = "hello\n"

// $ go test -bench=.
func BenchmarkWriteFile(b *testing.B) {

	for n := 0; n < b.N; n++ {

		f, err := os.Create(FILE_PATH)

		if err != nil {
			panic(err)
		}

		for i := 0; i < LINES_COUNT; i++ {

			f.WriteString(TEXT)
		}

		f.Close()
	}

}

/*
$ go test -bench=.
goos: linux
goarch: amd64
pkg: training.go/goapprofondi
BenchmarkWriteFile-4           4         351488127 ns/op
PASS
ok      training.go/goapprofondi        2.686s
*/

// $ go test -bench=.
func BenchmarkWriteFileBuffered(b *testing.B) {

	for n := 0; n < b.N; n++ {

		f, err := os.Create(FILE_PATH)

		if err != nil {
			panic(err)
		}

		// retourne un writer
		// bufio : buffer input/output
		w := bufio.NewWriter(f)

		for i := 0; i < LINES_COUNT; i++ {

			w.WriteString(TEXT)
		}

		w.Flush()
		f.Close()
	}

}

/*
$ go test -bench=.
goos: linux
goarch: amd64
pkg: training.go/goapprofondi
BenchmarkWriteFile-4                   3         336064310 ns/op
BenchmarkWriteFileBuffered-4         182           5533839 ns/op
PASS
ok      training.go/goapprofondi        3.001s
*/
