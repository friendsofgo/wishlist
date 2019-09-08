package wishlist

// Item represent our Item model
type Item struct {
	ID         string
	WishListID string
	Name       string
	Link       string
	Price      float64
	Priority   ItemPriority
	Status     ItemStatus
}

// ItemPriority uses priority to choose your level of desire to obtain this item
type ItemPriority int

const (
	// HighItemPriority represent the most high priority for your items
	HighItemPriority ItemPriority = 100
	// MidItemPriority represent the  middle priority for your items
	MidItemPriority ItemPriority = 50
	// LowItemPriority represent the lowest priority for your items
	LowItemPriority ItemPriority = 0
)

func (i ItemPriority) String() (priority string) {
	switch i {
	case HighItemPriority:
		priority = "high"
	case MidItemPriority:
		priority = "moderate"
	case LowItemPriority:
		priority = "low"
	}

	return
}

type ItemStatus int

const (
	// InactiveItem define inactive item status
	InactiveItem ItemStatus = iota
	// ActiveItem define active item status
	ActiveItem
)
