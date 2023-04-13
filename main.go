package main

import (
	"example.com/courier-service/model"
	"example.com/courier-service/service"
)

func main() {

	deliveryInfo := new(model.DeliveryInfo)

	model.InitializeCouponMap(deliveryInfo)
	service.AddCoupon(deliveryInfo, nil)
	service.GetCalculationMode(deliveryInfo, model.Undefined)
	service.GetInitialInput(deliveryInfo)
	service.GetPkgListInput(deliveryInfo)
	service.GetFinalInput(deliveryInfo)

	service.CalculateDeliveryCost(deliveryInfo)
	service.CalculateDeliveryTime(deliveryInfo)

}
