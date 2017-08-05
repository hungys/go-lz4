go-lz4
======

[![GoDoc](https://godoc.org/github.com/hungys/go-lz4?status.svg)](https://godoc.org/github.com/hungys/go-lz4)
[![Build Status](https://travis-ci.org/hungys/go-lz4.svg?branch=master)](https://travis-ci.org/hungys/go-lz4)

go-lz4 is Go bindings for [LZ4](https://github.com/lz4/lz4) compression library.

# Usage

```go
import "github.com/hungys/go-lz4"
```

Here are all the available APIs,

```
func CompressBound(size int) int
func CompressDefault(source, dest []byte) (int, error)
func CompressFast(source, dest []byte, acceleration int) (int, error)

func DecompressSafe(source, dest []byte) (int, error)
func DecompressFast(source, dest []byte, originalSize int) (int, error)
```

We also provide another sets of APIs for compressing or decompressing struct or slice. For struct, you should pass a **pointer**. But noted that for slice, you should pass **slice itself**, but not a pointer!

```
func CompressAnyDefault(source, dest interface{}, sourceSize, maxDestSize int) (int, error)
func CompressAnyFast(source, dest interface{}, sourceSize, maxDestSize, acceleration int) (int, error)

func DecompressAnySafe(source, dest interface{}, compressedSize, maxDecompressedSize int) (int, error)
func DecompressAnyFast(source, dest interface{}, originalSize int) (int, error)
```
