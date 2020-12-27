package searcher

import (
	datastorerPkg "pulley.com/shakesearch/pkg/datastore"
)

// Searcher specifies the methods required to do a successul search operation
// it allows various search implementations to be used without changing code.
type Searcher interface {
	// Load(string) error
	Search(string, datastorerPkg.Storer) []string
}