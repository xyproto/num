# num

[![GoDoc](https://godoc.org/github.com/xyproto/num?status.svg)](http://godoc.org/github.com/xyproto/num)

Go module for dealing with fractions, and a utility for dividing two numbers and returning the simplified fraction.

## Example use

    > frac 2/5
    ⅖
    > frac 0.8
    ⅘
    > frac 123
    123

Use only 100 iterations when creating a fraction that represents the given float:

    > frac -m 100 0.777777777
    10/13

## Installation

    go install github.com/xyproto/num/cmd/frac@latest

## General info

* License: BSD-3
