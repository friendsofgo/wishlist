package grpc

import (
	"context"

	wishgrpc "github.com/friendsofgo/wishlist/genproto/go"
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
) wishgrpc.WishListServiceServer {
	return &wishListHandler{creatingService: cS, addingService: aS, listingService: lS}
}

func (s wishListHandler) Create(ctx context.Context, req *wishgrpc.CreateWishListReq) (*wishgrpc.CreateWishListResp, error) {
	id, err := s.creatingService.Create(req.WishList.Name, wishlist.Status(req.WishList.Status))
	if err != nil {
		return nil, err
	}
	return &wishgrpc.CreateWishListResp{WishListId: id}, nil
}

func (s wishListHandler) Add(ctx context.Context, req *wishgrpc.AddItemReq) (*wishgrpc.AddItemResp, error) {
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
	return &wishgrpc.AddItemResp{ItemId: id}, nil
}

func (s wishListHandler) List(ctx context.Context, req *wishgrpc.ListWishListReq) (*wishgrpc.ListWishListResp, error) {
	items, err := s.listingService.ListItems(req.WishListId)
	if err != nil {
		return nil, err
	}
	return &wishgrpc.ListWishListResp{Items: mapSliceOfItems(items)}, nil
}

func mapSliceOfItems(domainItems []wishlist.Item) (grpcItems []*wishgrpc.Item) {
	for _, i := range domainItems {
		grpcItems = append(grpcItems, &wishgrpc.Item{
			Id:         i.ID,
			WishListId: i.WishListID,
			Name:       i.Name,
			Link:       i.Link,
			Price:      i.Price,
			Priority:   wishgrpc.Item_ItemPriority(i.Priority),
			Status:     wishgrpc.Item_ItemStatus(i.Status),
		})
	}
	return
}
