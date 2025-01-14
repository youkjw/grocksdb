package grocksdb

import "testing"

func TestSSTFileReader(t *testing.T) {
	path := "/tmp/test.sst"
	opts := NewDefaultOptions()

	reader := NewSSTFileReader(opts)
	defer reader.Destroy()
	if err := reader.Open(path); err != nil {
		t.Fatal(err)
	}

	readOpts := NewDefaultReadOptions()
	iter := reader.NewIterator(readOpts)
	defer iter.Close()

	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		t.Logf("key: %s", string(iter.Key().Data()))
		t.Logf("value: %s", string(iter.Value().Data()))
	}

	t.Log("TestSSTFileReader pass")
}
