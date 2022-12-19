// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cart.proto

package cart

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Cart service

func NewCartEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cart service

type CartService interface {
	AddCart(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*ResponseAdd, error)
	CleanCart(ctx context.Context, in *Clean, opts ...client.CallOption) (*Response, error)
	Incr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	Decr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	DeleteItemByID(ctx context.Context, in *CartID, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *CartFindAll, opts ...client.CallOption) (*CartAll, error)
}

type cartService struct {
	c    client.Client
	name string
}

func NewCartService(name string, c client.Client) CartService {
	return &cartService{
		c:    c,
		name: name,
	}
}

func (c *cartService) AddCart(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*ResponseAdd, error) {
	req := c.c.NewRequest(c.name, "Cart.AddCart", in)
	out := new(ResponseAdd)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) CleanCart(ctx context.Context, in *Clean, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.CleanCart", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) Incr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.Incr", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) Decr(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.Decr", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) DeleteItemByID(ctx context.Context, in *CartID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.DeleteItemByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) GetAll(ctx context.Context, in *CartFindAll, opts ...client.CallOption) (*CartAll, error) {
	req := c.c.NewRequest(c.name, "Cart.GetAll", in)
	out := new(CartAll)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cart service

type CartHandler interface {
	AddCart(context.Context, *CartInfo, *ResponseAdd) error
	CleanCart(context.Context, *Clean, *Response) error
	Incr(context.Context, *Item, *Response) error
	Decr(context.Context, *Item, *Response) error
	DeleteItemByID(context.Context, *CartID, *Response) error
	GetAll(context.Context, *CartFindAll, *CartAll) error
}

func RegisterCartHandler(s server.Server, hdlr CartHandler, opts ...server.HandlerOption) error {
	type cart interface {
		AddCart(ctx context.Context, in *CartInfo, out *ResponseAdd) error
		CleanCart(ctx context.Context, in *Clean, out *Response) error
		Incr(ctx context.Context, in *Item, out *Response) error
		Decr(ctx context.Context, in *Item, out *Response) error
		DeleteItemByID(ctx context.Context, in *CartID, out *Response) error
		GetAll(ctx context.Context, in *CartFindAll, out *CartAll) error
	}
	type Cart struct {
		cart
	}
	h := &cartHandler{hdlr}
	return s.Handle(s.NewHandler(&Cart{h}, opts...))
}

type cartHandler struct {
	CartHandler
}

func (h *cartHandler) AddCart(ctx context.Context, in *CartInfo, out *ResponseAdd) error {
	return h.CartHandler.AddCart(ctx, in, out)
}

func (h *cartHandler) CleanCart(ctx context.Context, in *Clean, out *Response) error {
	return h.CartHandler.CleanCart(ctx, in, out)
}

func (h *cartHandler) Incr(ctx context.Context, in *Item, out *Response) error {
	return h.CartHandler.Incr(ctx, in, out)
}

func (h *cartHandler) Decr(ctx context.Context, in *Item, out *Response) error {
	return h.CartHandler.Decr(ctx, in, out)
}

func (h *cartHandler) DeleteItemByID(ctx context.Context, in *CartID, out *Response) error {
	return h.CartHandler.DeleteItemByID(ctx, in, out)
}

func (h *cartHandler) GetAll(ctx context.Context, in *CartFindAll, out *CartAll) error {
	return h.CartHandler.GetAll(ctx, in, out)
}
