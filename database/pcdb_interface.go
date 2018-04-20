package database

// PaulCacheDatabase is the standard for a cache database
type PaulCacheDatabase interface {
	AddRegion(string) error
	RemoveRegion(string) error
	GetRegion(string) (*Region, error)
}
