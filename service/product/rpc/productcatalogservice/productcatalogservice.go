// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: product.proto

package productcatalogservice

import (
	"context"

	"gomall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateProductReq   = product.CreateProductReq
	CreateProductResp  = product.CreateProductResp
	DecreaseStockReq   = product.DecreaseStockReq
	DecreaseStockResp  = product.DecreaseStockResp
	DeleteProductReq   = product.DeleteProductReq
	DeleteProductResp  = product.DeleteProductResp
	GetProductReq      = product.GetProductReq
	GetProductResp     = product.GetProductResp
	ListProductsReq    = product.ListProductsReq
	ListProductsResp   = product.ListProductsResp
	Product            = product.Product
	SearchProductsReq  = product.SearchProductsReq
	SearchProductsResp = product.SearchProductsResp
	StockItem          = product.StockItem
	UpdateProductReq   = product.UpdateProductReq
	UpdateProductResp  = product.UpdateProductResp

	ProductCatalogService interface {
		CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
		UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error)
		DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error)
		ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error)
		GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error)
		SearchProducts(ctx context.Context, in *SearchProductsReq, opts ...grpc.CallOption) (*SearchProductsResp, error)
		DecreaseStock(ctx context.Context, in *DecreaseStockReq, opts ...grpc.CallOption) (*DecreaseStockResp, error)
		DecreaseStockRevert(ctx context.Context, in *DecreaseStockReq, opts ...grpc.CallOption) (*DecreaseStockResp, error)
	}

	defaultProductCatalogService struct {
		cli zrpc.Client
	}
)

func NewProductCatalogService(cli zrpc.Client) ProductCatalogService {
	return &defaultProductCatalogService{
		cli: cli,
	}
}

func (m *defaultProductCatalogService) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.CreateProduct(ctx, in, opts...)
}

func (m *defaultProductCatalogService) UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.UpdateProduct(ctx, in, opts...)
}

func (m *defaultProductCatalogService) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.DeleteProduct(ctx, in, opts...)
}

func (m *defaultProductCatalogService) ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.ListProducts(ctx, in, opts...)
}

func (m *defaultProductCatalogService) GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.GetProduct(ctx, in, opts...)
}

func (m *defaultProductCatalogService) SearchProducts(ctx context.Context, in *SearchProductsReq, opts ...grpc.CallOption) (*SearchProductsResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.SearchProducts(ctx, in, opts...)
}

func (m *defaultProductCatalogService) DecreaseStock(ctx context.Context, in *DecreaseStockReq, opts ...grpc.CallOption) (*DecreaseStockResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.DecreaseStock(ctx, in, opts...)
}

func (m *defaultProductCatalogService) DecreaseStockRevert(ctx context.Context, in *DecreaseStockReq, opts ...grpc.CallOption) (*DecreaseStockResp, error) {
	client := product.NewProductCatalogServiceClient(m.cli.Conn())
	return client.DecreaseStockRevert(ctx, in, opts...)
}
