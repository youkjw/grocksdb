package grocksdb

import (
	"bytes"
	"testing"
)

func TestSSTFileWriter(t *testing.T) {
	path := "/tmp/test.sst"
	envOpts := NewDefaultEnvOptions()
	opts := NewDefaultOptions()

	writer := NewSSTFileWriterWithComparator(envOpts, opts, NewComparator("order", func(a, b []byte) int {
		return bytes.Compare(a, b)
	}))
	defer writer.Destroy()
	if err := writer.Open(path); err != nil {
		t.Fatal(err)
	}

	if err := writer.Put([]byte("aaa"), []byte("aaa")); err != nil {
		t.Fatal(err)
	}
	if err := writer.Put([]byte("bbb"), []byte("bbb")); err != nil {
		t.Fatal(err)
	}
	if err := writer.Put([]byte("ccc"), []byte("ccc")); err != nil {
		t.Fatal(err)
	}
	if err := writer.Put([]byte("ddd"), []byte("ddd")); err != nil {
		t.Fatal(err)
	}

	if err := writer.Finish(); err != nil {
		t.Fatal(err)
	}

	t.Log("TestSSTFileWriter pass")
}
