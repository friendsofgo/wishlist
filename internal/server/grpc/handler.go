package grpc

import (
	"context"

	grpc "github.com/friendsofgo/wishlist/genproto/go"
	wishlist "github.com/friendsofgo/wishlist/internal"
	"github.com/friendsofgo/wishlist/internal/adding"
	"github.com/friendsofgo/wishlist/internal/creating"
	"github.com/friendsofgo/wishlist/internal/listing"
)

type wishListHandler struct {
	creatingService creating.Service
	addingService   adding.Service
	listingService  listing.Service
}

// NewWishListServer provides WishList gRPC operations
func NewWishListServer(
	cS creating.Service,
	aS adding.Service,
	lS listing.Service,
) grpc.WishListServiceServer {
	return &wishListHandler{creatingService: cS, addingService: aS, listingService: lS}
}

func (s wishListHandler) Create(ctx context.Context, req *grpc.CreateWishListReq) (*grpc.CreateWishListResp, error) {
	id, err := s.creatingService.Create(req.WishList.Name, wishlist.Status(req.WishList.Status))
	if err != nil {
		return nil, err
	}
	return &grpc.CreateWishListResp{WishListId: id}, nil
}

func (s wishListHandler) Add(ctx context.Context, req *grpc.AddItemReq) (*grpc.AddItemResp, error) {
	id, err := s.addingService.AddItem(
		req.Item.WishListId,
		req.Item.Name,
		req.Item.Link,
		req.Item.Price,
		wishlist.ItemPriority(req.Item.Priority),
		wishlist.ItemStatus(req.Item.Status),
	)
	if err != nil {
		return nil, err
	}
	return &grpc.AddItemResp{ItemId: id}, nil
}

func (s wishListHandler) List(ctx context.Context, req *grpc.ListWishListReq) (*grpc.ListWishListResp, error) {
	items, err := s.listingService.ListItems(req.WishListId)
	if err != nil {
		return nil, err
	}
	return &grpc.ListWishListResp{Items: mapSliceOfItems(items)}, nil
}

func mapSliceOfItems(domainItems []wishlist.Item) (grpcItems []*grpc.Item) {
	for _, i := range domainItems {
		grpcItems = append(grpcItems, &grpc.Item{
			Id:         i.ID,
			WishListId: i.WishListID,
			Name:       i.Name,
			Link:       i.Link,
			Price:      i.Price,
			Priority:   grpc.Item_ItemPriority(i.Priority),
			Status:     grpc.Item_ItemStatus(i.Status),
		})
	}
	return
}
