package logic

import (
	"context"
	"gomall/service/checkout/api/internal/svc"
	"gomall/service/checkout/api/internal/types"
	"gomall/service/checkout/rpc/types/checkout"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckoutLogic {
	return &CheckoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckoutLogic) Checkout(req *types.CheckoutRequest) (resp *types.CheckoutResponse, err error) {
	var cartItems []*checkout.CartItem
	for _, item := range req.CartItems {
		cartItems = append(cartItems, &checkout.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	res, err := l.svcCtx.CheckoutRpc.Checkout(l.ctx, &checkout.CheckoutReq{
		UserId:    uint32(l.ctx.Value("userId").(int64)),
		Currency:  req.Currency,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &checkout.Address{
			StreetAddress: req.UserAddress.StreetAddress,
			City:          req.UserAddress.City,
			State:         req.UserAddress.State,
			Country:       req.UserAddress.Country,
			ZipCode:       req.UserAddress.ZipCode,
		},
		CreditCard: &checkout.CreditCardInfo{
			CreditCardNumber:          req.CardInfo.CreditCardNumber,
			CreditCardCvv:             req.CardInfo.CreditCardCVV,
			CreditCardExpirationYear:  req.CardInfo.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CardInfo.CreditCardExpirationMonth,
		},
		CartItems: cartItems,
	})
	if err != nil {
		return nil, err
	}

	return &types.CheckoutResponse{
		OrderId:       res.OrderId,
		TransactionId: res.TransactionId,
	}, nil
}
