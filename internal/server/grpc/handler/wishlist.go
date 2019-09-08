package handler

import (
	"context"
	"errors"
	"log"

	"github.com/friendsofgo/wishlist/internal/net/grpc"
)

type wishListHandler struct {
}

// NewWishList provides wishList gRPC operations
func NewWishList() grpc.WishListServiceServer {
	return &wishListHandler{}
}

func (s *wishListHandler) List(*grpc.ListWishListReq, grpc.WishListService_ListServer) error {
	return errors.New("method not implemented yet")
}

func (s *wishListHandler) Create(context.Context, *grpc.CreateWishListReq) (*grpc.CreateWishListResp, error) {
	log.Println("create handler")
	return nil, errors.New("method not implemented yet")
}

func (s *wishListHandler) Update(context.Context, *grpc.UpdateWishListReq) (*grpc.UpdateWishListResp, error) {
	return nil, errors.New("method not implemented yet")
}
