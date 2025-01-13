// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: cart.proto

package cartservice

import (
	"context"

	"gomall/service/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddItemReq    = cart.AddItemReq
	AddItemResp   = cart.AddItemResp
	Cart          = cart.Cart
	CartItem      = cart.CartItem
	EmptyCartReq  = cart.EmptyCartReq
	EmptyCartResp = cart.EmptyCartResp
	GetCartReq    = cart.GetCartReq
	GetCartResp   = cart.GetCartResp

	CartService interface {
		AddItem(ctx context.Context, in *AddItemReq, opts ...grpc.CallOption) (*AddItemResp, error)
		GetCart(ctx context.Context, in *GetCartReq, opts ...grpc.CallOption) (*GetCartResp, error)
		EmptyCart(ctx context.Context, in *EmptyCartReq, opts ...grpc.CallOption) (*EmptyCartResp, error)
	}

	defaultCartService struct {
		cli zrpc.Client
	}
)

func NewCartService(cli zrpc.Client) CartService {
	return &defaultCartService{
		cli: cli,
	}
}

func (m *defaultCartService) AddItem(ctx context.Context, in *AddItemReq, opts ...grpc.CallOption) (*AddItemResp, error) {
	client := cart.NewCartServiceClient(m.cli.Conn())
	return client.AddItem(ctx, in, opts...)
}

func (m *defaultCartService) GetCart(ctx context.Context, in *GetCartReq, opts ...grpc.CallOption) (*GetCartResp, error) {
	client := cart.NewCartServiceClient(m.cli.Conn())
	return client.GetCart(ctx, in, opts...)
}

func (m *defaultCartService) EmptyCart(ctx context.Context, in *EmptyCartReq, opts ...grpc.CallOption) (*EmptyCartResp, error) {
	client := cart.NewCartServiceClient(m.cli.Conn())
	return client.EmptyCart(ctx, in, opts...)
}
