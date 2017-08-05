go-lz4
======

[![GoDoc](https://godoc.org/github.com/hungys/go-lz4?status.svg)](https://godoc.org/github.com/hungys/go-lz4)
[![Build Status](https://travis-ci.org/hungys/go-lz4.svg?branch=master)](https://travis-ci.org/hungys/go-lz4)

go-lz4 is a Go binding for [LZ4](https://github.com/lz4/lz4) compression library.

Here are all available APIs,

```
func CompressBound(size int) int
func CompressDefault(source, dest []byte) (int, error)
func CompressFast(source, dest []byte, acceleration int) (int, error)

func DecompressSafe(source, dest []byte) (int, error)
func DecompressFast(source, dest []byte, originalSize int) (int, error)
```
