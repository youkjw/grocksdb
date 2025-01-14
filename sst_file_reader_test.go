package grocksdb

import "testing"

func TestSSTFileReader(t *testing.T) {
	path := "/tmp/test.sst"
	envOpts := NewDefaultEnvOptions()
	opts := NewDefaultOptions()

	reader := NewSSTFileReader(envOpts, opts)
	reader.Open(path)
}
