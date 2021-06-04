// go test main_test.go
// go test ./... -v

// https://tutorialedge.net/golang/improving-your-tests-with-testify-go/

// https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318

// The first thing to do in order to get up and running with the testify package is to install it.
// Now, if you are using Go Modules then this will just be a case of calling go test ... after importing
// the package at the top of one of your *_test.go files.

// On an older version of Go you can get this package by typing:
// go get github.com/stretchr/testify

// This works -not so much
// go mod init github.com/stretchr/testify
// go mod tidy
// go get github.com/stretchr/testify/assert

// export GOPROXY=https://gocenter.io
// go get github.com/stretchr/testify/assert

// go1.15.13 version

package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Without import "github.com/stretchr/testify/assert"
// We can then try run this simple test by calling go test ./... -v
// passing in the -v flag to ensure we can see a more verbose output.
// func testAdd_3(t *testing.T) {
//     if add(2, 2) != 4 {
//         t.Error("Expected 2 + 2 to equal 4")
//     }
// }


func testAdd(t *testing.T) {
    assert.Equal(t, add(2, 2), 4)
}


// func testAdd(t *testing.T) {
//     assert := assert.New(t)
//
//     var tests = []struct {
//         input1    int
//         input2    int
//         expected  int
//     }{
//         {2, 2, 4},
//         {-1, 1, 0},
//         {0, 2, 2},
//         {-5, -3, -8},
//     }
//
//     for _, test := range tests {
//         assert.Equal(add(test.input1, test.input2), test.expected)
//     }
// }
