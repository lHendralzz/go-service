package order

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-service/model"
	"go-service/stdlib/auth"
	x "go-service/stdlib/error"
	"net/http"
	"time"

	"github.com/palantir/stacktrace"
)

func (o *orderService) CheckoutOrder(ctx context.Context, param model.CheckoutOrderRequest) (int, error) {
	tx, err := o.orderRepository.BeginTransaction()
	if err != nil {
		return 0, stacktrace.Propagate(err, "Failed Begin Transaction CheckoutOrder")
	}

	orderID, err := o.orderRepository.AddOrderWithTx(tx, param)
	if err != nil {
		errRollback := tx.Rollback().Error
		if errRollback != nil {
			o.logger.Error(errRollback)
		}
		return 0, err
	}

	o.logger.Debug("success Create Order With ID: ", orderID)

	key := fmt.Sprintf("checkout:shop:%d", param.ShopID)
	o.logger.Debug("Acquire Lock : ", key)

	acquire, err := o.redis.AcquireLock(ctx, key, time.Second*time.Duration(o.opt.LockDuration))
	if err != nil {
		return 0, stacktrace.PropagateWithCode(err, x.ErrorRedisLock, "Failed AcquireLock")
	}

	defer func() {
		o.redis.ReleaseLock(ctx, key)
		o.logger.Debug("release Lock : ", key)
	}()

	if !acquire {
		return 0, stacktrace.NewErrorWithCode(x.ErrorLockedOrder, "Cannot Checkout Order Locked Shop")
	}

	// TODO : SELECT id, name, price, stock FROM product p WHERE id in (1,2,3);
	// mapped to map[id]stock remainingStock

	for _, v := range param.Products {

		// TODO check remaininStock[v.ProductID] < v.Quantity
		// send request to product
		err := o.sendRequestToAddStock(model.AddStockProductRequest{
			ProductID: v.ProductID,
			Quantity:  -v.Quantity,
		})
		if err != nil {
			errRollback := tx.Rollback().Error
			if errRollback != nil {
				o.logger.Error(errRollback)
			}
			return 0, stacktrace.PropagateWithCode(err, x.ErrorReduceProduct, "Failed SendRequestToReduceStock")
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return 0, stacktrace.Propagate(err, "Failed Commit Transaction CheckoutOrder")
	}

	return orderID, nil
}

func (o *orderService) sendRequestToAddStock(request model.AddStockProductRequest) error {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return stacktrace.Propagate(err, "Failed Marshal")
	}

	// TODO :localhost change to product service
	url := "http://localhost:8080" + fmt.Sprintf("/product/%d/add-stock", request.ProductID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return stacktrace.Propagate(err, "Failed Create New Request")
	}

	// TODO : fix this from generate token
	token, err := auth.GenerateToken("admin", 1, "SECRET_JWT_KEY")
	if err != nil {
		return stacktrace.Propagate(err, "Failed Generate Token")
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return stacktrace.Propagate(err, "Failed Create Send Request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return stacktrace.NewErrorWithCode(x.ErrorReduceProduct, "Failed SendRequestToReduceStock Status = %d", resp.StatusCode)
	}

	// Send response to the original client
	return nil
}

// ReleaseOrderFromCheckoutStatus function to return stock from order that has not been process from x duration
func (o *orderService) ReleaseOrderFromCheckoutStatus(x time.Duration) error {
	now := time.Now().Add(-x)
	o.logger.Info(fmt.Sprintf("Get Order That has Been made before %s and has status checkout (1)", now.String()))
	// Get Order That has Been made before now and has status checkout (1)

	// 	Loop All Expired Order
	// 		release Order
	// 		Lock shop
	// 		get order details by order ID
	// 			loop all product in order details
	// 			sendRequestToAddStock(model.AddStockProductRequest{
	//				ProductID: v.ProductID,
	//				Quantity:  -v.Quantity,
	// 			})
	// 		Unlock shop
	return nil
}
