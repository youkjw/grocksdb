package grocksdb

import (
	"fmt"
	"testing"
)

func TestSSTFileWriter(t *testing.T) {
	path := "/tmp/test.sst"
	envOpts := NewDefaultEnvOptions()
	opts := NewDefaultOptions()

	writer := NewSSTFileWriter(envOpts, opts)
	defer writer.Destroy()
	if err := writer.Open(path); err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 1e3; i++ {
		k, v := fmt.Sprintf("test_%d", i), fmt.Sprintf("value_%d", i)
		writer.Add([]byte(k), []byte(v))
	}
	if err := writer.Finish(); err != nil {
		t.Fatal(err)
	}

	t.Log("TestSSTFileWriter pass")
}
