package creating

import wishlist "github.com/friendsofgo/wishlist/internal"

type Service interface {
	Create(name string, status wishlist.Status) (string, error)
}

type service struct {
	repository wishlist.Repository
}

func (s service) Create(name string, status wishlist.Status) (string, error) {
	id := wishlist.NewULID()
	w := wishlist.WishList{ID: id, Name: name, Status: status}
	err := s.repository.Store(w)
	if err != nil {
		return "", err
	}
	return id, nil
}

func NewService(repository wishlist.Repository) Service {
	return service{repository: repository}
}
