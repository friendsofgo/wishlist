package wishlist

// WishList represent our WishList model
type WishList struct {
	ID     string
	Name   string
	Status Status
}

// Status type to define the wish lists status
type Status int

const (
	// Inactive define inactive wish list status
	Inactive Status = iota
	// Active define active wish list status
	Active
)

// Repository provide operations above repository
type Repository interface {
	Store(WishList) error
	AddItem(string, Item) error
	FetchItems(string) ([]Item, error)
}
