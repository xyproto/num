# num

[![Build Status](https://travis-ci.org/xyproto/num.svg?branch=master)](https://travis-ci.org/xyproto/num) [![GoDoc](https://godoc.org/github.com/xyproto/num?status.svg)](http://godoc.org/github.com/xyproto/num)

Utility for dividing two numbers and returning the simplified fraction. Will use unicode, if possible.

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

## Go module

Also includes a Go module that provides the same functionality.

## Installation

Development version:

    go get -u github.com/xyproto/num/cmd/frac

With make:

    make && sudo make install

