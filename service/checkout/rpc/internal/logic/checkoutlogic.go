package logic

import (
	"context"
	"fmt"
	"gomall/service/order/rpc/types/order"
	"gomall/service/payment/rpc/types/payment"
	"gomall/service/product/model"
	"gomall/service/product/rpc/types/product"
	"google.golang.org/grpc/status"
	"strconv"

	"gomall/service/checkout/rpc/internal/svc"
	"gomall/service/checkout/rpc/types/checkout"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckoutLogic {
	return &CheckoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckoutLogic) Checkout(in *checkout.CheckoutReq) (*checkout.CheckoutResp, error) {
	// 先检查库存的量是否足够
	productIds := make([]uint32, len(in.CartItems))
	for i, item := range in.CartItems {
		productIds[i] = item.ProductId
	}
	var productModels []*model.ProductModel
	err := l.svcCtx.DB.Find(&productModels, productIds).Error
	if err != nil {
		return nil, status.Errorf(500, "fail to query product stock, %v", err.Error())
	}

	// 计算总金额
	totalPrice := 0.0
	var orderItems []*order.OrderItem
	for _, item := range in.CartItems {
		for _, product_ := range productModels {
			if item.ProductId == product_.Id {
				if int64(item.Quantity) > product_.Stock {
					return nil, status.Errorf(500, "product %d stock is not enough", item.ProductId)
				}
				totalPrice += product_.Price * float64(item.Quantity)
				orderItems = append(orderItems, &order.OrderItem{
					Item: &order.CartItem{
						ProductId: item.ProductId,
						Quantity:  item.Quantity,
					},
					Cost: float32(product_.Price),
				})
			}
		}
	}

	var items []*product.StockItem
	// 扣减库存,但是后续操作如果失败，需要回滚库存
	for _, item := range in.CartItems {
		for i := 0; i < len(productModels); i++ {
			if item.ProductId == productModels[i].Id {
				productModels[i].Stock -= int64(item.Quantity)

				items = append(items, &product.StockItem{
					ProductId: item.ProductId,
					Quantity:  int64(item.Quantity),
				})
			}
		}

	}
	err = l.svcCtx.DB.Save(&productModels).Error
	if err != nil {
		return nil, status.Errorf(500, "failed to deduct product stock: %v", err.Error())
	}

	orderRes, err := l.svcCtx.OrderRpc.PlaceOrder(l.ctx, &order.PlaceOrderReq{
		UserId:       in.UserId,
		UserCurrency: in.Currency,
		Email:        in.Email,
		Address: &order.Address{
			StreetAddress: in.Address.StreetAddress,
			City:          in.Address.City,
			State:         in.Address.State,
			Country:       in.Address.Country,
			ZipCode:       in.Address.ZipCode,
		},
		OrderItems: orderItems,
	})
	if err != nil {
		logx.Errorf("fail to place order,try to rollback database. %v", err.Error())
		// 回滚库存
		err1 := rollbackStock(l, in.CartItems, productModels)
		if err1 != nil {
			logx.Errorf("critical error during rollback: %v", err1.Error())
			return nil, status.Errorf(500, "fail to rollback product stock, %v", err1.Error())
		}
		return nil, status.Errorf(500, "fail to place order, %v", err.Error())
	}

	chargeRes, err := l.svcCtx.PaymentRpc.Charge(l.ctx, &payment.ChargeReq{
		Amount: float32(totalPrice),
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          in.CreditCard.CreditCardNumber,
			CreditCardCvv:             in.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  in.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: in.CreditCard.CreditCardExpirationMonth,
		},
		OrderId: orderRes.Order.OrderId,
		UserId:  in.UserId,
	})
	logx.Infof("chargeRes: %v", chargeRes)
	if err != nil || chargeRes.GetSuccess() == "false" {
		logx.Error("fail to charge, try to rollback database")
		// 回滚库存
		err1 := rollbackStock(l, in.CartItems, productModels)
		if err1 != nil {
			logx.Errorf("critical error during rollback: %v", err1.Error())
			return nil, status.Errorf(500, "fail to rollback product stock, %v", err1.Error())
		}
		_, err = l.svcCtx.OrderRpc.MarkOrderCanceled(l.ctx, &order.MarkOrderCanceledReq{
			OrderId: orderRes.Order.OrderId,
			UserId:  in.UserId,
		})
		if err != nil {
			logx.Errorf("fail to mark order canceled, try to rollback database. %v", err.Error())
			// TODO 通过消息队列，发送消息到消息队列，让订单服务去执行该取消订单操作
		}

		return nil, status.Error(500, "fail to charge")
	}
	_, err = l.svcCtx.OrderRpc.MarkOrderPaid(l.ctx, &order.MarkOrderPaidReq{
		OrderId: orderRes.Order.OrderId,
		UserId:  in.UserId,
	})
	if err != nil {
		logx.Errorf("fail to mark order paid, try to rollback database. %v", err.Error())
		// TODO 通过消息队列，发送消息到消息队列，让订单服务去执行标记订单已支付
	}

	return &checkout.CheckoutResp{OrderId: strconv.Itoa(int(orderRes.Order.OrderId)), TransactionId: strconv.FormatUint(chargeRes.TransactionId, 10)}, nil
}

// 库存回滚逻辑
func rollbackStock(l *CheckoutLogic, cartItems []*checkout.CartItem, productModels []*model.ProductModel) error {
	for _, item := range cartItems {
		for i := 0; i < len(productModels); i++ {
			if item.ProductId == productModels[i].Id {
				productModels[i].Stock += int64(item.Quantity)
			}
		}
	}
	err := l.svcCtx.DB.Save(&productModels).Error
	if err != nil {
		logx.Errorf("failed to rollback stock: %v", err.Error())
		return fmt.Errorf("failed to rollback stock: %w", err)
	}
	return nil
}
