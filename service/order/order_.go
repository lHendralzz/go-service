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

func (o *orderService) CheckoutOrder(ctx context.Context, param model.CheckoutOrderRequest) error {
	tx, err := o.orderRepository.BeginTransaction()
	if err != nil {
		return stacktrace.Propagate(err, "Failed Begin Transaction CheckoutOrder")
	}

	orderID, err := o.orderRepository.AddOrderWithTx(tx, param)
	if err != nil {
		errRollback := tx.Rollback().Error
		if errRollback != nil {
			o.logger.Error(errRollback)
		}
		return err
	}

	o.logger.Info("success Create Order With ID: ", orderID)

	key := fmt.Sprintf("checkout:shop:%d", param.ShopID)
	o.logger.Info("Acquire Lock : ", key)

	acquire, err := o.redis.AcquireLock(ctx, key, time.Second*time.Duration(o.opt.LockDuration))
	if err != nil {
		return stacktrace.PropagateWithCode(err, x.ErrorRedisLock, "Failed AcquireLock")
	}

	defer func() {
		o.redis.ReleaseLock(ctx, key)
		o.logger.Info("release Lock : ", key)
	}()

	if !acquire {
		return stacktrace.NewErrorWithCode(x.ErrorLockedOrder, "Cannot Checkout Order Locked Shop")
	}

	for _, v := range param.Products {
		// send request to product
		err := o.SendRequestToReduceStock(model.AddStockProductRequest{
			ProductID: v.ProductID,
			Quantity:  -v.Quantity,
		})
		if err != nil {
			errRollback := tx.Rollback().Error
			if errRollback != nil {
				o.logger.Error(errRollback)
			}
			return stacktrace.PropagateWithCode(err, x.ErrorReduceProduct, "Failed SendRequestToReduceStock")
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return stacktrace.Propagate(err, "Failed Commit Transaction CheckoutOrder")
	}

	return nil
}

func (o orderService) SendRequestToReduceStock(request model.AddStockProductRequest) error {
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
