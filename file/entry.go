package file

// Entry of the file
type Entry struct {
	Address string
	Domain  string
}

// NewEntry constructor
func NewEntry(domain string, address string) *Entry {
	return &Entry{
		Address: address,
		Domain:  domain,
	}
}
