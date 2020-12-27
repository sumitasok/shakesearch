package datastore

// Storer interface for hosting data in file/ in memory / in database
type Storer interface {
	QueryIDs(string) []int
	GetByID(int) string
}