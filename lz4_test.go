package lz4

import (
	"crypto/rand"
	"testing"
	"unsafe"
)

type TestStruct struct {
	a int32
	b int32
	c [1024]int32
}

func generateTestStruct() TestStruct {
	ts := TestStruct{}
	ts.a = 111
	ts.b = 222
	for i := 0; i < 1024; i++ {
		ts.c[i] = int32(i / 100)
	}
	return ts
}

func areByteSlicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCompressDefaultDecompressSafe(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, CompressBound(4096))
	size, err := CompressDefault(data, compressed)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	decompressed := make([]byte, 4096)
	if _, err := DecompressSafe(compressed[:size], decompressed); err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}

	if !areByteSlicesEqual(data, decompressed) {
		t.Fatal("Incorrect decompression output")
	}
}

func TestCompressDefaultWithInsufficientBuffer(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, 1)
	_, err := CompressDefault(data, compressed)
	if err == nil {
		t.Fatalf("Compression should be failed due to insufficient buffer")
	}
}

func TestDecompressSafeWithInsufficientBuffer(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, CompressBound(4096))
	size, err := CompressDefault(data, compressed)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	decompressed := make([]byte, 1)
	if _, err := DecompressSafe(compressed[:size], decompressed); err == nil {
		t.Fatalf("Decompression should be failed due to insufficient buffer")
	}
}

func TestCompressFastDecompressFast(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, CompressBound(4096))
	size, err := CompressFast(data, compressed, 9)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	decompressed := make([]byte, 4096)
	if _, err := DecompressFast(compressed[:size], decompressed, len(decompressed)); err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}

	if !areByteSlicesEqual(data, decompressed) {
		t.Fatal("Incorrect decompression output")
	}
}

func TestCompressFastWithInsufficientBuffer(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, 1)
	_, err := CompressFast(data, compressed, 9)
	if err == nil {
		t.Fatalf("Compression should be failed due to insufficient buffer")
	}
}

func TestDecompressFastWithInsufficientBuffer(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, CompressBound(4096))
	size, err := CompressFast(data, compressed, 9)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	decompressed := make([]byte, 1)
	if _, err := DecompressFast(compressed[:size], decompressed, len(decompressed)); err == nil {
		t.Fatalf("Decompression should be failed due to insufficient buffer")
	}
}

func TestCompressAnyDefaultDecompressAnySafe(t *testing.T) {
	data := generateTestStruct()
	dataSize := int(unsafe.Sizeof(data))

	compressed := make([]byte, CompressBound(dataSize))
	size, err := CompressAnyDefault(&data, compressed, dataSize, len(compressed))
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	decompressed := TestStruct{}
	if _, err := DecompressAnySafe(compressed[:size], &decompressed, size, dataSize); err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}

	if decompressed != data {
		t.Fatal("Incorrect decompression output")
	}
}

func TestCompressAnyDefaultWithInsufficientBuffer(t *testing.T) {
	data := generateTestStruct()
	dataSize := int(unsafe.Sizeof(data))

	compressed := make([]byte, 1)
	_, err := CompressAnyDefault(&data, compressed, dataSize, len(compressed))
	if err == nil {
		t.Fatalf("Compression should be failed due to insufficient buffer")
	}
}

func TestDecompressAnySafeWithInsufficientBuffer(t *testing.T) {
	data := generateTestStruct()
	dataSize := int(unsafe.Sizeof(data))

	compressed := make([]byte, CompressBound(dataSize))
	size, err := CompressAnyDefault(&data, compressed, dataSize, len(compressed))
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	var decompressed int32
	if _, err := DecompressAnySafe(compressed[:size], &decompressed, size, int(unsafe.Sizeof(decompressed))); err == nil {
		t.Fatalf("Decompression should be failed due to insufficient buffer")
	}
}

func TestCompressAnyFastDecompressAnyFast(t *testing.T) {
	data := generateTestStruct()
	dataSize := int(unsafe.Sizeof(data))

	compressed := make([]byte, CompressBound(dataSize))
	size, err := CompressAnyFast(&data, compressed, dataSize, len(compressed), 9)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	decompressed := TestStruct{}
	if _, err := DecompressAnyFast(compressed[:size], &decompressed, dataSize); err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}

	if decompressed != data {
		t.Fatal("Incorrect decompression output")
	}
}

func TestCompressAnyFastWithInsufficientBuffer(t *testing.T) {
	data := generateTestStruct()
	dataSize := int(unsafe.Sizeof(data))

	compressed := make([]byte, 1)
	_, err := CompressAnyFast(&data, compressed, dataSize, len(compressed), 9)
	if err == nil {
		t.Fatalf("Compression should be failed due to insufficient buffer")
	}
}

func TestDecompressAnyFastWithInsufficientBuffer(t *testing.T) {
	data := generateTestStruct()
	dataSize := int(unsafe.Sizeof(data))

	compressed := make([]byte, CompressBound(dataSize))
	size, err := CompressAnyFast(&data, compressed, dataSize, len(compressed), 9)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	var decompressed int32
	if _, err := DecompressAnyFast(compressed[:size], &decompressed, int(unsafe.Sizeof(decompressed))); err == nil {
		t.Fatalf("Decompression should be failed due to insufficient buffer")
	}
}
