package shakesearch

import (
	"fmt"
	"index/suffixarray"
	// "io/ioutil"

	datastorerPkg "pulley.com/shakesearch/pkg/datastore"
	searcherPkg "pulley.com/shakesearch/pkg/searcher"
)

// we ensure ShakeSearcher implements expected interface
var _ searcherPkg.Searcher = Searcher{}

// Searcher implements nterfac Searcher and uses an algorithm we call ShakeSearch
type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

// func (s *Searcher) Load(filename string) error {
// 	dat, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return fmt.Errorf("Load: %w", err)
// 	}
// 	s.CompleteWorks = string(dat)
// 	s.SuffixArray = suffixarray.New(dat)
// 	return nil
// }

// Search ...
func (s Searcher) Search(query string, store datastorerPkg.Storer) []string {
	idxs := store.QueryIDs(query)

	results := []string{}
	for _, idx := range idxs {
		// N+1
		results = append(results, store.GetByID(idx))
	}
	fmt.Println(results)
	return results
}
