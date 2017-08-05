package lz4

// #cgo CFLAGS: -O3
// #include "lz4.h"
import "C"
import (
	"errors"
	"unsafe"
)

func byteSliceToCharPointer(in []byte) *C.char {
	if len(in) == 0 {
		return (*C.char)(unsafe.Pointer(nil))
	}
	return (*C.char)(unsafe.Pointer(&in[0]))
}

func CompressDefault(source, dest []byte) (int, error) {
	ret := int(C.LZ4_compress_default(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(len(source)), C.int(len(dest))))
	if ret < 0 {
		return ret, errors.New("Insufficient destination buffer")
	}

	return ret, nil
}

func CompressBound(size int) int {
	return int(C.LZ4_compressBound(C.int(size)))
}

func CompressFast(source, dest []byte, acceleration int) (int, error) {
	ret := int(C.LZ4_compress_fast(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(len(source)), C.int(len(dest)),
		C.int(acceleration)))
	if ret < 0 {
		return ret, errors.New("Insufficient destination buffer")
	}

	return ret, nil
}

func DecompressSafe(source, dest []byte) (int, error) {
	ret := int(C.LZ4_decompress_safe(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(len(source)), C.int(len(dest))))
	if ret < 0 {
		return ret, errors.New("Malformed LZ4 source or insufficient destination buffer")
	}

	return ret, nil
}

func DecompressFast(source, dest []byte, originalSize int) (int, error) {
	ret := int(C.LZ4_decompress_fast(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(originalSize)))
	if ret < 0 {
		return ret, errors.New("Malformed LZ4 source")
	}

	return ret, nil
}
