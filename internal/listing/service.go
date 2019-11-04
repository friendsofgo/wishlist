package listing

import wishlist "github.com/friendsofgo/wishlist/internal"

type Service interface {
	ListItems(string) ([]wishlist.Item, error)
}

type service struct {
	repository wishlist.Repository
}

func (s service) ListItems(wishListID string) ([]wishlist.Item, error) {
	return s.repository.FetchItems(wishListID)
}

func NewService(repository wishlist.Repository) Service {
	return service{repository: repository}
}
