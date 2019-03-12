package file

// Entry of the file
type Entry struct {
	Address string
	Domain  string
}

// NewEntry constructor
func NewEntry(address string, domain string) *Entry {
	return &Entry{
		Address: address,
		Domain:  domain,
	}
}
