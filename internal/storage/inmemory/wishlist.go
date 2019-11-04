package inmemory

import (
	"errors"

	wishlist "github.com/friendsofgo/wishlist/internal"
)

var (
	WishListNotFound      = errors.New("wish list not found")
	WishListAlreadyExists = errors.New("wish list already exists")
)

type inmemoryWishListRepo struct {
	wishLists map[string]wishlist.WishList
	items     map[string][]wishlist.Item
}

func NewInMemoryWishListRepository() wishlist.Repository {
	return inmemoryWishListRepo{}
}

func (r inmemoryWishListRepo) Store(w wishlist.WishList) error {
	if _, ok := r.wishLists[w.ID]; ok {
		return WishListAlreadyExists
	}
	r.wishLists[w.ID] = w
	return nil
}

func (r inmemoryWishListRepo) AddItem(ID string, item wishlist.Item) error {
	if _, ok := r.wishLists[ID]; !ok {
		return WishListNotFound
	}
	r.items[ID] = append(r.items[ID], item)
	return nil
}

func (r inmemoryWishListRepo) FetchItems(ID string) ([]wishlist.Item, error) {
	if _, ok := r.wishLists[ID]; !ok {
		return []wishlist.Item{}, WishListNotFound
	}
	return r.items[ID], nil
}
