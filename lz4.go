package lz4

// #cgo CFLAGS: -O3
// #include "lz4.h"
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

func byteSliceToCharPointer(b []byte) *C.char {
	if len(b) == 0 {
		return (*C.char)(unsafe.Pointer(nil))
	}
	return (*C.char)(unsafe.Pointer(&b[0]))
}

func interfaceToCharPointer(i interface{}) *C.char {
	if i == nil {
		return (*C.char)(unsafe.Pointer(nil))
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Slice {
		return (*C.char)(unsafe.Pointer(v.Index(0).UnsafeAddr()))
	} else if v.Kind() == reflect.Ptr {
		return (*C.char)(unsafe.Pointer(v.Pointer()))
	}

	return (*C.char)(unsafe.Pointer(nil))
}

// CompressBound returns the maximum size that LZ4 compression may output in a "worst case" scenario (input data not compressible).
func CompressBound(size int) int {
	return int(C.LZ4_compressBound(C.int(size)))
}

// CompressDefault compresses buffer "source" into already allocated "dest" buffer.
// Compression is guaranteed to succeed if size of "dest" >= CompressBound(size of "src")
// The function returns the number of bytes written into buffer "dest".
// If the function cannot compress "source" into a more limited "dest" budget,
// compression stops immediately, and the function result is zero.
func CompressDefault(source, dest []byte) (int, error) {
	ret := int(C.LZ4_compress_default(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(len(source)), C.int(len(dest))))
	if ret == 0 {
		return ret, errors.New("Insufficient destination buffer")
	}

	return ret, nil
}

// CompressFast works the same as CompressDefault, but allows to select an "acceleration" factor.
// The larger the acceleration value, the faster the algorithm, but also the lesser the compression.
// An acceleration value of "1" is the same as regular CompressDefault()
func CompressFast(source, dest []byte, acceleration int) (int, error) {
	ret := int(C.LZ4_compress_fast(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(len(source)), C.int(len(dest)),
		C.int(acceleration)))
	if ret == 0 {
		return ret, errors.New("Insufficient destination buffer")
	}

	return ret, nil
}

// DecompressSafe decompresses buffer "source" into already allocated "dest" buffer.
// The function returns the number of bytes written into buffer "dest".
// If destination buffer is not large enough, decoding will stop and output an error code (<0).
// If the source stream is detected malformed, the function will stop decoding and return a negative result.
func DecompressSafe(source, dest []byte) (int, error) {
	ret := int(C.LZ4_decompress_safe(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(len(source)), C.int(len(dest))))
	if ret < 0 {
		return ret, errors.New("Malformed LZ4 source or insufficient destination buffer")
	}

	return ret, nil
}

// DecompressFast fully respect memory boundaries for properly formed compressed data.
// It is a bit faster than DecompressSafe.
// However, it does not provide any protection against intentionally modified data stream (malicious input).
func DecompressFast(source, dest []byte, originalSize int) (int, error) {
	ret := int(C.LZ4_decompress_fast(byteSliceToCharPointer(source),
		byteSliceToCharPointer(dest), C.int(originalSize)))
	if ret < 0 {
		return ret, errors.New("Malformed LZ4 source")
	}

	return ret, nil
}

// CompressAnyDefault works the same as CompressDefault, but you can pass pointer of struct or slice as a buffer.
// Also, the size of buffer should be specified explicitly.
func CompressAnyDefault(source, dest interface{}, sourceSize, maxDestSize int) (int, error) {
	ret := int(C.LZ4_compress_default(interfaceToCharPointer(source),
		interfaceToCharPointer(dest), C.int(sourceSize), C.int(maxDestSize)))
	if ret == 0 {
		return ret, errors.New("Insufficient destination buffer")
	}

	return ret, nil
}

// CompressAnyFast works the same as CompressFast, but you can pass pointer of struct or slice as a buffer.
// Also, the size of buffer should be specified explicitly.
func CompressAnyFast(source, dest interface{}, sourceSize, maxDestSize, acceleration int) (int, error) {
	ret := int(C.LZ4_compress_fast(interfaceToCharPointer(source),
		interfaceToCharPointer(dest), C.int(sourceSize), C.int(maxDestSize),
		C.int(acceleration)))
	if ret == 0 {
		return ret, errors.New("Insufficient destination buffer")
	}

	return ret, nil
}

// DecompressAnySafe works the same as DecompressSafe, but you can pass pointer of struct or slice as a buffer.
// Also, the size of buffer should be specified explicitly.
func DecompressAnySafe(source, dest interface{}, compressedSize, maxDecompressedSize int) (int, error) {
	ret := int(C.LZ4_decompress_safe(interfaceToCharPointer(source),
		interfaceToCharPointer(dest), C.int(compressedSize), C.int(maxDecompressedSize)))
	if ret < 0 {
		return ret, errors.New("Malformed LZ4 source or insufficient destination buffer")
	}

	return ret, nil
}

// DecompressAnyFast works the same as DecompressFast, but you can pass pointer of struct or slice as a buffer.
func DecompressAnyFast(source, dest interface{}, originalSize int) (int, error) {
	ret := int(C.LZ4_decompress_fast(interfaceToCharPointer(source),
		interfaceToCharPointer(dest), C.int(originalSize)))
	if ret < 0 {
		return ret, errors.New("Malformed LZ4 source")
	}

	return ret, nil
}
