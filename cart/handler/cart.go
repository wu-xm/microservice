package handler

import (
	"cart/domain/service"
	pb "cart/proto"
	"context"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c *Cart) AddCart(ctx context.Context, info *pb.CartInfo, add *pb.ResponseAdd) error {
	//TODO implement me
	panic("implement me")
}

func (c *Cart) CleanCart(ctx context.Context, clean *pb.Clean, response *pb.Response) error {
	//TODO implement me
	panic("implement me")
}

func (c *Cart) Incr(ctx context.Context, item *pb.Item, response *pb.Response) error {
	//TODO implement me
	panic("implement me")
}

func (c *Cart) Decr(ctx context.Context, item *pb.Item, response *pb.Response) error {
	//TODO implement me
	panic("implement me")
}

func (c *Cart) DeleteItemByID(ctx context.Context, id *pb.CartID, response *pb.Response) error {
	//TODO implement me
	panic("implement me")
}

func (c *Cart) GetAll(ctx context.Context, all *pb.CartFindAll, all2 *pb.CartAll) error {
	//TODO implement me
	panic("implement me")
}
