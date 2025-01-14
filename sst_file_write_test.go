package grocksdb

import (
	"bytes"
	"fmt"
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

	for i := 0; i <= 1000; i++ {
		k, v := fmt.Sprintf("%d", i), fmt.Sprintf("%d", i)
		if err := writer.Put([]byte(k), []byte(v)); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Finish(); err != nil {
		t.Fatal(err)
	}

	t.Log("TestSSTFileWriter pass")
}
