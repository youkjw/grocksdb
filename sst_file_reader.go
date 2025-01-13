package grocksdb

// #include <stdlib.h>
// #include "rocksdb/c.h"
import "C"

import (
	"unsafe"
)

type SSTFileReader struct {
	c *C.rocksdb_sstfilereader_t
}

// NewSSTFileReader creates an SSTFileReader object.
func NewSSTFileReader(opts *EnvOptions, dbOpts *Options) *SSTFileReader {
	c := C.rocksdb_sstfilewriter_create(opts.c, dbOpts.c)
	return &SSTFileReader{c: c}
}

// Open prepares SstFileWriter to write into file located at "path".
func (w *SSTFileReader) Open(path string) (err error) {
	var (
		cErr  *C.char
		cPath = C.CString(path)
	)

	C.rocksdb_sstfilereader_open(w.c, cPath, &cErr)
	err = fromCError(cErr)

	C.free(unsafe.Pointer(cPath))
	return
}

func (w *SSTFileReader) NewIterator(dbOpts *ReadOptions) (iter *Iterator) {
	iter = C.rocksdb_sstfilereader_iterator(w.c, dbOpts.c)
	return
}

// Verifies whether there is corruption in this table.
func (w *SSTFileReader) VerifyChecksum(dbOpts *ReadOptions) (err error) {
	var (
		cErr *C.char
	)
	C.rocksdb_sstfilereader_verifychecksum(w.c, &cErr)
	err = fromCError(cErr)
	return
}

// Destroy destroys the SSTFileReader object.
func (w *SSTFileReader) Destroy() {
	C.rocksdb_sstfilereader_destroy(w.c)
	w.c = nil
}
