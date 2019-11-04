package adding

import wishlist "github.com/friendsofgo/wishlist/internal"

type Service interface {
	AddItem(wishListID string, name, link string, price float64, priority wishlist.ItemPriority, status wishlist.ItemStatus) (string, error)
}

type service struct {
	repository wishlist.Repository
}

func (s service) AddItem(wishListID string, name, link string, price float64, priority wishlist.ItemPriority, status wishlist.ItemStatus) (string, error) {
	id := wishlist.NewULID()
	item := wishlist.Item{ID: id, WishListID: wishListID, Name: name, Link: link, Price: price, Priority: priority, Status: status}
	err := s.repository.AddItem(wishListID, item)
	return id, err
}

func NewService(repository wishlist.Repository) Service {
	return service{repository: repository}
}
