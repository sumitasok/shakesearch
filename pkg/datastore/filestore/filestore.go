package filestore

import (
	"fmt"
	"index/suffixarray"
	"io/ioutil"
)

// FileStore data store to hold ShakeSpear code
type FileStore struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

// QueryIDs does a basic query on the Store / FileStore in this case.
func (fS FileStore) QueryIDs(query string) []int {
	idxs := fS.SuffixArray.Lookup([]byte(query), -1)

	return idxs
}

// GetByID returns the element at the location provided
func (fS FileStore) GetByID(idx int) string {
	return fS.CompleteWorks[idx-250:idx+250]
}

// NewFileStore return data base from file provided.
func NewFileStore(filename string) (FileStore, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return FileStore{}, fmt.Errorf("Load: %w", err)
	}

	return FileStore{
		CompleteWorks: string(dat),
		SuffixArray: suffixarray.New(dat),
	}, nil
}