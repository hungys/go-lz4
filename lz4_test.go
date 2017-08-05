package lz4

import (
	"crypto/rand"
	"testing"
)

func isByteSliceEqual(a, b []byte) bool {
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

	if !isByteSliceEqual(data, decompressed) {
		t.Fatal("Incorrect decompression output")
	}
}

func TestCompressDefaultWithInsufficientBuffer(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, 1)
	_, err := CompressDefault(data, compressed)
	if err == nil {
		t.Fatalf("Compression should be failed")
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
		t.Fatalf("Decompression should be failed")
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
	if _, err := DecompressFast(compressed[:size], decompressed, 4096); err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}

	if !isByteSliceEqual(data, decompressed) {
		t.Fatal("Incorrect decompression output")
	}
}

func TestCompressFastWithInsufficientBuffer(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, 1)
	_, err := CompressFast(data, compressed, 9)
	if err == nil {
		t.Fatalf("Compression should be failed")
	}
}

func TestDecompressFastWithInsufficientBuffer(t *testing.T) {
	compressed := make([]byte, CompressBound(4096))
	rand.Read(compressed)

	decompressed := make([]byte, 1)
	if _, err := DecompressFast(compressed, decompressed, 4096); err == nil {
		t.Fatalf("Decompression should be failed")
	}
}
